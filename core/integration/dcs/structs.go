/*
 * === This file is part of ALICE O² ===
 *
 * Copyright 2021-2023 CERN and copyright holders of ALICE O².
 * Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * In applying this license CERN does not waive the privileges and
 * immunities granted to it by virtue of its status as an
 * Intergovernmental Organization or submit itself to any jurisdiction.
 */

package dcs

import (
	"fmt"
	"strings"

	dcspb "github.com/AliceO2Group/Control/core/integration/dcs/protos"
)

func (d DCSDetectors) ToStringSlice() (sslice []string) {
	if d == nil {
		return
	}
	sslice = make([]string, len(d))
	if len(d) == 0 {
		return
	}

	for i, det := range d {
		sslice[i] = det.String()
	}
	return
}

func (d DCSDetectors) EcsDetectorsSlice() (sslice []string) {
	if d == nil {
		return
	}
	sslice = make([]string, len(d))
	if len(d) == 0 {
		return
	}

	for i, det := range d {
		sslice[i] = dcsToEcsDetector(det)
	}
	return
}

func (dsm DCSDetectorStatesMap) makeDetectorsByStateMap() map[dcspb.DetectorState]DCSDetectors {
	detectorsByState := make(map[dcspb.DetectorState]DCSDetectors)
	for det, detState := range dsm {
		if _, ok := detectorsByState[detState]; !ok {
			detectorsByState[detState] = make([]dcspb.Detector, 0)
		}
		detectorsByState[detState] = append(detectorsByState[detState], det)
	}
	return detectorsByState
}

// Returns true if the provided detectors are either all in conditionState or in NULL_STATE
func (dsm DCSDetectorStatesMap) compatibleWithDCSOperation(conditionState dcspb.DetectorState) (bool, error) {
	detectorsByState := dsm.makeDetectorsByStateMap()

	detectorsInConditionState, thereAreDetectorsInConditionState := detectorsByState[conditionState]
	detectorsInNullState, thereAreDetectorsInNullState := detectorsByState[dcspb.DetectorState_NULL_STATE]

	if thereAreDetectorsInConditionState && (len(detectorsInConditionState) == len(dsm)) {
		// all detectors are in conditionState
		return true, nil
	} else if thereAreDetectorsInConditionState && thereAreDetectorsInNullState && (len(detectorsInConditionState)+len(detectorsInNullState) == len(dsm)) {
		// all detectors are either in conditionState or in NULL_STATE
		return true, fmt.Errorf("detectors %s are in NULL_STATE", strings.Join(detectorsByState[dcspb.DetectorState_NULL_STATE].ToStringSlice(), ", "))
	} else if thereAreDetectorsInNullState && (len(detectorsInNullState) == len(dsm)) {
		// all detectors are in NULL_STATE
		return true, fmt.Errorf("all detectors are in NULL_STATE")
	} else {
		// there are detectors in other states incompatible with conditionState
		reportByState := make([]string, 0)
		for state, detectors := range detectorsByState {
			if state == conditionState {
				continue
			}
			reportByState = append(reportByState,
				fmt.Sprintf("%s in %s", strings.Join(detectors.ToStringSlice(), ", "), state.String()))
		}
		return false, fmt.Errorf("detectors are in incompatible states: %v", strings.Join(reportByState, "; "))
	}
}
