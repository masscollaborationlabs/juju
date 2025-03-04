// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/charmrevisionupdater (interfaces: Application,CharmhubRefreshClient,Model,State)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/mocks.go github.com/juju/juju/apiserver/facades/controller/charmrevisionupdater Application,CharmhubRefreshClient,Model,State
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/charm/v11"
	charmrevisionupdater "github.com/juju/juju/apiserver/facades/controller/charmrevisionupdater"
	charmhub "github.com/juju/juju/charmhub"
	transport "github.com/juju/juju/charmhub/transport"
	cloud "github.com/juju/juju/cloud"
	metrics "github.com/juju/juju/core/charm/metrics"
	config "github.com/juju/juju/environs/config"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// ApplicationTag mocks base method.
func (m *MockApplication) ApplicationTag() names.ApplicationTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationTag")
	ret0, _ := ret[0].(names.ApplicationTag)
	return ret0
}

// ApplicationTag indicates an expected call of ApplicationTag.
func (mr *MockApplicationMockRecorder) ApplicationTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationTag", reflect.TypeOf((*MockApplication)(nil).ApplicationTag))
}

// CharmOrigin mocks base method.
func (m *MockApplication) CharmOrigin() *state.CharmOrigin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmOrigin")
	ret0, _ := ret[0].(*state.CharmOrigin)
	return ret0
}

// CharmOrigin indicates an expected call of CharmOrigin.
func (mr *MockApplicationMockRecorder) CharmOrigin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmOrigin", reflect.TypeOf((*MockApplication)(nil).CharmOrigin))
}

// CharmURL mocks base method.
func (m *MockApplication) CharmURL() (*string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CharmURL indicates an expected call of CharmURL.
func (mr *MockApplicationMockRecorder) CharmURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockApplication)(nil).CharmURL))
}

// UnitCount mocks base method.
func (m *MockApplication) UnitCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// UnitCount indicates an expected call of UnitCount.
func (mr *MockApplicationMockRecorder) UnitCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitCount", reflect.TypeOf((*MockApplication)(nil).UnitCount))
}

// MockCharmhubRefreshClient is a mock of CharmhubRefreshClient interface.
type MockCharmhubRefreshClient struct {
	ctrl     *gomock.Controller
	recorder *MockCharmhubRefreshClientMockRecorder
}

// MockCharmhubRefreshClientMockRecorder is the mock recorder for MockCharmhubRefreshClient.
type MockCharmhubRefreshClientMockRecorder struct {
	mock *MockCharmhubRefreshClient
}

// NewMockCharmhubRefreshClient creates a new mock instance.
func NewMockCharmhubRefreshClient(ctrl *gomock.Controller) *MockCharmhubRefreshClient {
	mock := &MockCharmhubRefreshClient{ctrl: ctrl}
	mock.recorder = &MockCharmhubRefreshClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmhubRefreshClient) EXPECT() *MockCharmhubRefreshClientMockRecorder {
	return m.recorder
}

// RefreshWithMetricsOnly mocks base method.
func (m *MockCharmhubRefreshClient) RefreshWithMetricsOnly(arg0 context.Context, arg1 map[metrics.MetricKey]map[metrics.MetricKey]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWithMetricsOnly", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshWithMetricsOnly indicates an expected call of RefreshWithMetricsOnly.
func (mr *MockCharmhubRefreshClientMockRecorder) RefreshWithMetricsOnly(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWithMetricsOnly", reflect.TypeOf((*MockCharmhubRefreshClient)(nil).RefreshWithMetricsOnly), arg0, arg1)
}

// RefreshWithRequestMetrics mocks base method.
func (m *MockCharmhubRefreshClient) RefreshWithRequestMetrics(arg0 context.Context, arg1 charmhub.RefreshConfig, arg2 map[metrics.MetricKey]map[metrics.MetricKey]string) ([]transport.RefreshResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWithRequestMetrics", arg0, arg1, arg2)
	ret0, _ := ret[0].([]transport.RefreshResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshWithRequestMetrics indicates an expected call of RefreshWithRequestMetrics.
func (mr *MockCharmhubRefreshClientMockRecorder) RefreshWithRequestMetrics(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWithRequestMetrics", reflect.TypeOf((*MockCharmhubRefreshClient)(nil).RefreshWithRequestMetrics), arg0, arg1, arg2)
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// CloudName mocks base method.
func (m *MockModel) CloudName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudName indicates an expected call of CloudName.
func (mr *MockModelMockRecorder) CloudName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudName", reflect.TypeOf((*MockModel)(nil).CloudName))
}

// CloudRegion mocks base method.
func (m *MockModel) CloudRegion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudRegion")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudRegion indicates an expected call of CloudRegion.
func (mr *MockModelMockRecorder) CloudRegion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudRegion", reflect.TypeOf((*MockModel)(nil).CloudRegion))
}

// Config mocks base method.
func (m *MockModel) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockModelMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockModel)(nil).Config))
}

// IsControllerModel mocks base method.
func (m *MockModel) IsControllerModel() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsControllerModel")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsControllerModel indicates an expected call of IsControllerModel.
func (mr *MockModelMockRecorder) IsControllerModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsControllerModel", reflect.TypeOf((*MockModel)(nil).IsControllerModel))
}

// Metrics mocks base method.
func (m *MockModel) Metrics() (state.ModelMetrics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metrics")
	ret0, _ := ret[0].(state.ModelMetrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metrics indicates an expected call of Metrics.
func (mr *MockModelMockRecorder) Metrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metrics", reflect.TypeOf((*MockModel)(nil).Metrics))
}

// ModelTag mocks base method.
func (m *MockModel) ModelTag() names.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockModelMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockModel)(nil).ModelTag))
}

// UUID mocks base method.
func (m *MockModel) UUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UUID indicates an expected call of UUID.
func (mr *MockModelMockRecorder) UUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUID", reflect.TypeOf((*MockModel)(nil).UUID))
}

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// AddCharmPlaceholder mocks base method.
func (m *MockState) AddCharmPlaceholder(arg0 *charm.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharmPlaceholder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCharmPlaceholder indicates an expected call of AddCharmPlaceholder.
func (mr *MockStateMockRecorder) AddCharmPlaceholder(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharmPlaceholder", reflect.TypeOf((*MockState)(nil).AddCharmPlaceholder), arg0)
}

// AliveRelationKeys mocks base method.
func (m *MockState) AliveRelationKeys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AliveRelationKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AliveRelationKeys indicates an expected call of AliveRelationKeys.
func (mr *MockStateMockRecorder) AliveRelationKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AliveRelationKeys", reflect.TypeOf((*MockState)(nil).AliveRelationKeys))
}

// AllApplications mocks base method.
func (m *MockState) AllApplications() ([]charmrevisionupdater.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllApplications")
	ret0, _ := ret[0].([]charmrevisionupdater.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllApplications indicates an expected call of AllApplications.
func (mr *MockStateMockRecorder) AllApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllApplications", reflect.TypeOf((*MockState)(nil).AllApplications))
}

// Charm mocks base method.
func (m *MockState) Charm(arg0 string) (*state.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(*state.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockStateMockRecorder) Charm(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockState)(nil).Charm), arg0)
}

// Cloud mocks base method.
func (m *MockState) Cloud(arg0 string) (cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud", arg0)
	ret0, _ := ret[0].(cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cloud indicates an expected call of Cloud.
func (mr *MockStateMockRecorder) Cloud(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockState)(nil).Cloud), arg0)
}

// ControllerUUID mocks base method.
func (m *MockState) ControllerUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerUUID indicates an expected call of ControllerUUID.
func (mr *MockStateMockRecorder) ControllerUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerUUID", reflect.TypeOf((*MockState)(nil).ControllerUUID))
}

// Model mocks base method.
func (m *MockState) Model() (charmrevisionupdater.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(charmrevisionupdater.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockStateMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockState)(nil).Model))
}

// Resources mocks base method.
func (m *MockState) Resources() state.Resources {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resources")
	ret0, _ := ret[0].(state.Resources)
	return ret0
}

// Resources indicates an expected call of Resources.
func (mr *MockStateMockRecorder) Resources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resources", reflect.TypeOf((*MockState)(nil).Resources))
}
