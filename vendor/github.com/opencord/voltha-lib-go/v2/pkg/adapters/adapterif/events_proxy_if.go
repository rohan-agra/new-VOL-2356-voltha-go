/*
 * Copyright 2018-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package adapterif

import (
	"github.com/opencord/voltha-protos/v2/go/voltha"
)

// EventProxy interface for eventproxy
type EventProxy interface {
	SendDeviceEvent(deviceEvent *voltha.DeviceEvent, category EventCategory,
		subCategory EventSubCategory, raisedTs int64) error
	SendKpiEvent(id string, deviceEvent *voltha.KpiEvent2, category EventCategory,
		subCategory EventSubCategory, raisedTs int64) error
}

const (
	EventTypeVersion = "0.1"
)

type (
	EventType        = voltha.EventType_EventType
	EventCategory    = voltha.EventCategory_EventCategory
	EventSubCategory = voltha.EventSubCategory_EventSubCategory
)
