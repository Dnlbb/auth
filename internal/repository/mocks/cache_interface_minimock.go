// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Dnlbb/auth/internal/repository/repoInterface.CacheInterface -o cache_interface_minimock.go -n CacheInterfaceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/gojuno/minimock/v3"
)

// CacheInterfaceMock implements mm_repoInterface.CacheInterface
type CacheInterfaceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, id int64, user models.User) (err error)
	funcCreateOrigin    string
	inspectFuncCreate   func(ctx context.Context, id int64, user models.User)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mCacheInterfaceMockCreate

	funcGet          func(ctx context.Context, params models.GetUserParams) (up1 *models.User, err error)
	funcGetOrigin    string
	inspectFuncGet   func(ctx context.Context, params models.GetUserParams)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mCacheInterfaceMockGet
}

// NewCacheInterfaceMock returns a mock for mm_repoInterface.CacheInterface
func NewCacheInterfaceMock(t minimock.Tester) *CacheInterfaceMock {
	m := &CacheInterfaceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mCacheInterfaceMockCreate{mock: m}
	m.CreateMock.callArgs = []*CacheInterfaceMockCreateParams{}

	m.GetMock = mCacheInterfaceMockGet{mock: m}
	m.GetMock.callArgs = []*CacheInterfaceMockGetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mCacheInterfaceMockCreate struct {
	optional           bool
	mock               *CacheInterfaceMock
	defaultExpectation *CacheInterfaceMockCreateExpectation
	expectations       []*CacheInterfaceMockCreateExpectation

	callArgs []*CacheInterfaceMockCreateParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// CacheInterfaceMockCreateExpectation specifies expectation struct of the CacheInterface.Create
type CacheInterfaceMockCreateExpectation struct {
	mock               *CacheInterfaceMock
	params             *CacheInterfaceMockCreateParams
	paramPtrs          *CacheInterfaceMockCreateParamPtrs
	expectationOrigins CacheInterfaceMockCreateExpectationOrigins
	results            *CacheInterfaceMockCreateResults
	returnOrigin       string
	Counter            uint64
}

// CacheInterfaceMockCreateParams contains parameters of the CacheInterface.Create
type CacheInterfaceMockCreateParams struct {
	ctx  context.Context
	id   int64
	user models.User
}

// CacheInterfaceMockCreateParamPtrs contains pointers to parameters of the CacheInterface.Create
type CacheInterfaceMockCreateParamPtrs struct {
	ctx  *context.Context
	id   *int64
	user *models.User
}

// CacheInterfaceMockCreateResults contains results of the CacheInterface.Create
type CacheInterfaceMockCreateResults struct {
	err error
}

// CacheInterfaceMockCreateOrigins contains origins of expectations of the CacheInterface.Create
type CacheInterfaceMockCreateExpectationOrigins struct {
	origin     string
	originCtx  string
	originId   string
	originUser string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreate *mCacheInterfaceMockCreate) Optional() *mCacheInterfaceMockCreate {
	mmCreate.optional = true
	return mmCreate
}

// Expect sets up expected params for CacheInterface.Create
func (mmCreate *mCacheInterfaceMockCreate) Expect(ctx context.Context, id int64, user models.User) *mCacheInterfaceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &CacheInterfaceMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.paramPtrs != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by ExpectParams functions")
	}

	mmCreate.defaultExpectation.params = &CacheInterfaceMockCreateParams{ctx, id, user}
	mmCreate.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// ExpectCtxParam1 sets up expected param ctx for CacheInterface.Create
func (mmCreate *mCacheInterfaceMockCreate) ExpectCtxParam1(ctx context.Context) *mCacheInterfaceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &CacheInterfaceMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &CacheInterfaceMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreate.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreate
}

// ExpectIdParam2 sets up expected param id for CacheInterface.Create
func (mmCreate *mCacheInterfaceMockCreate) ExpectIdParam2(id int64) *mCacheInterfaceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &CacheInterfaceMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &CacheInterfaceMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.id = &id
	mmCreate.defaultExpectation.expectationOrigins.originId = minimock.CallerInfo(1)

	return mmCreate
}

// ExpectUserParam3 sets up expected param user for CacheInterface.Create
func (mmCreate *mCacheInterfaceMockCreate) ExpectUserParam3(user models.User) *mCacheInterfaceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &CacheInterfaceMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &CacheInterfaceMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.user = &user
	mmCreate.defaultExpectation.expectationOrigins.originUser = minimock.CallerInfo(1)

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the CacheInterface.Create
func (mmCreate *mCacheInterfaceMockCreate) Inspect(f func(ctx context.Context, id int64, user models.User)) *mCacheInterfaceMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for CacheInterfaceMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by CacheInterface.Create
func (mmCreate *mCacheInterfaceMockCreate) Return(err error) *CacheInterfaceMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &CacheInterfaceMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &CacheInterfaceMockCreateResults{err}
	mmCreate.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreate.mock
}

// Set uses given function f to mock the CacheInterface.Create method
func (mmCreate *mCacheInterfaceMockCreate) Set(f func(ctx context.Context, id int64, user models.User) (err error)) *CacheInterfaceMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the CacheInterface.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the CacheInterface.Create method")
	}

	mmCreate.mock.funcCreate = f
	mmCreate.mock.funcCreateOrigin = minimock.CallerInfo(1)
	return mmCreate.mock
}

// When sets expectation for the CacheInterface.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mCacheInterfaceMockCreate) When(ctx context.Context, id int64, user models.User) *CacheInterfaceMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CacheInterfaceMock.Create mock is already set by Set")
	}

	expectation := &CacheInterfaceMockCreateExpectation{
		mock:               mmCreate.mock,
		params:             &CacheInterfaceMockCreateParams{ctx, id, user},
		expectationOrigins: CacheInterfaceMockCreateExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up CacheInterface.Create return parameters for the expectation previously defined by the When method
func (e *CacheInterfaceMockCreateExpectation) Then(err error) *CacheInterfaceMock {
	e.results = &CacheInterfaceMockCreateResults{err}
	return e.mock
}

// Times sets number of times CacheInterface.Create should be invoked
func (mmCreate *mCacheInterfaceMockCreate) Times(n uint64) *mCacheInterfaceMockCreate {
	if n == 0 {
		mmCreate.mock.t.Fatalf("Times of CacheInterfaceMock.Create mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreate.expectedInvocations, n)
	mmCreate.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreate
}

func (mmCreate *mCacheInterfaceMockCreate) invocationsDone() bool {
	if len(mmCreate.expectations) == 0 && mmCreate.defaultExpectation == nil && mmCreate.mock.funcCreate == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreate.mock.afterCreateCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreate.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Create implements mm_repoInterface.CacheInterface
func (mmCreate *CacheInterfaceMock) Create(ctx context.Context, id int64, user models.User) (err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	mmCreate.t.Helper()

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, id, user)
	}

	mm_params := CacheInterfaceMockCreateParams{ctx, id, user}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_want_ptrs := mmCreate.CreateMock.defaultExpectation.paramPtrs

		mm_got := CacheInterfaceMockCreateParams{ctx, id, user}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreate.t.Errorf("CacheInterfaceMock.Create got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreate.CreateMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmCreate.t.Errorf("CacheInterfaceMock.Create got unexpected parameter id, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreate.CreateMock.defaultExpectation.expectationOrigins.originId, *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

			if mm_want_ptrs.user != nil && !minimock.Equal(*mm_want_ptrs.user, mm_got.user) {
				mmCreate.t.Errorf("CacheInterfaceMock.Create got unexpected parameter user, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreate.CreateMock.defaultExpectation.expectationOrigins.originUser, *mm_want_ptrs.user, mm_got.user, minimock.Diff(*mm_want_ptrs.user, mm_got.user))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("CacheInterfaceMock.Create got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreate.CreateMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the CacheInterfaceMock.Create")
		}
		return (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, id, user)
	}
	mmCreate.t.Fatalf("Unexpected call to CacheInterfaceMock.Create. %v %v %v", ctx, id, user)
	return
}

// CreateAfterCounter returns a count of finished CacheInterfaceMock.Create invocations
func (mmCreate *CacheInterfaceMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of CacheInterfaceMock.Create invocations
func (mmCreate *CacheInterfaceMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to CacheInterfaceMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mCacheInterfaceMockCreate) Calls() []*CacheInterfaceMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*CacheInterfaceMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *CacheInterfaceMock) MinimockCreateDone() bool {
	if m.CreateMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateMock.invocationsDone()
}

// MinimockCreateInspect logs each unmet expectation
func (m *CacheInterfaceMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CacheInterfaceMock.Create at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateCounter := mm_atomic.LoadUint64(&m.afterCreateCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && afterCreateCounter < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to CacheInterfaceMock.Create at\n%s", m.CreateMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to CacheInterfaceMock.Create at\n%s with params: %#v", m.CreateMock.defaultExpectation.expectationOrigins.origin, *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && afterCreateCounter < 1 {
		m.t.Errorf("Expected call to CacheInterfaceMock.Create at\n%s", m.funcCreateOrigin)
	}

	if !m.CreateMock.invocationsDone() && afterCreateCounter > 0 {
		m.t.Errorf("Expected %d calls to CacheInterfaceMock.Create at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateMock.expectedInvocations), m.CreateMock.expectedInvocationsOrigin, afterCreateCounter)
	}
}

type mCacheInterfaceMockGet struct {
	optional           bool
	mock               *CacheInterfaceMock
	defaultExpectation *CacheInterfaceMockGetExpectation
	expectations       []*CacheInterfaceMockGetExpectation

	callArgs []*CacheInterfaceMockGetParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// CacheInterfaceMockGetExpectation specifies expectation struct of the CacheInterface.Get
type CacheInterfaceMockGetExpectation struct {
	mock               *CacheInterfaceMock
	params             *CacheInterfaceMockGetParams
	paramPtrs          *CacheInterfaceMockGetParamPtrs
	expectationOrigins CacheInterfaceMockGetExpectationOrigins
	results            *CacheInterfaceMockGetResults
	returnOrigin       string
	Counter            uint64
}

// CacheInterfaceMockGetParams contains parameters of the CacheInterface.Get
type CacheInterfaceMockGetParams struct {
	ctx    context.Context
	params models.GetUserParams
}

// CacheInterfaceMockGetParamPtrs contains pointers to parameters of the CacheInterface.Get
type CacheInterfaceMockGetParamPtrs struct {
	ctx    *context.Context
	params *models.GetUserParams
}

// CacheInterfaceMockGetResults contains results of the CacheInterface.Get
type CacheInterfaceMockGetResults struct {
	up1 *models.User
	err error
}

// CacheInterfaceMockGetOrigins contains origins of expectations of the CacheInterface.Get
type CacheInterfaceMockGetExpectationOrigins struct {
	origin       string
	originCtx    string
	originParams string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGet *mCacheInterfaceMockGet) Optional() *mCacheInterfaceMockGet {
	mmGet.optional = true
	return mmGet
}

// Expect sets up expected params for CacheInterface.Get
func (mmGet *mCacheInterfaceMockGet) Expect(ctx context.Context, params models.GetUserParams) *mCacheInterfaceMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &CacheInterfaceMockGetExpectation{}
	}

	if mmGet.defaultExpectation.paramPtrs != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by ExpectParams functions")
	}

	mmGet.defaultExpectation.params = &CacheInterfaceMockGetParams{ctx, params}
	mmGet.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// ExpectCtxParam1 sets up expected param ctx for CacheInterface.Get
func (mmGet *mCacheInterfaceMockGet) ExpectCtxParam1(ctx context.Context) *mCacheInterfaceMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &CacheInterfaceMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &CacheInterfaceMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.ctx = &ctx
	mmGet.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGet
}

// ExpectParamsParam2 sets up expected param params for CacheInterface.Get
func (mmGet *mCacheInterfaceMockGet) ExpectParamsParam2(params models.GetUserParams) *mCacheInterfaceMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &CacheInterfaceMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &CacheInterfaceMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.params = &params
	mmGet.defaultExpectation.expectationOrigins.originParams = minimock.CallerInfo(1)

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the CacheInterface.Get
func (mmGet *mCacheInterfaceMockGet) Inspect(f func(ctx context.Context, params models.GetUserParams)) *mCacheInterfaceMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for CacheInterfaceMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by CacheInterface.Get
func (mmGet *mCacheInterfaceMockGet) Return(up1 *models.User, err error) *CacheInterfaceMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &CacheInterfaceMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &CacheInterfaceMockGetResults{up1, err}
	mmGet.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGet.mock
}

// Set uses given function f to mock the CacheInterface.Get method
func (mmGet *mCacheInterfaceMockGet) Set(f func(ctx context.Context, params models.GetUserParams) (up1 *models.User, err error)) *CacheInterfaceMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the CacheInterface.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the CacheInterface.Get method")
	}

	mmGet.mock.funcGet = f
	mmGet.mock.funcGetOrigin = minimock.CallerInfo(1)
	return mmGet.mock
}

// When sets expectation for the CacheInterface.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mCacheInterfaceMockGet) When(ctx context.Context, params models.GetUserParams) *CacheInterfaceMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CacheInterfaceMock.Get mock is already set by Set")
	}

	expectation := &CacheInterfaceMockGetExpectation{
		mock:               mmGet.mock,
		params:             &CacheInterfaceMockGetParams{ctx, params},
		expectationOrigins: CacheInterfaceMockGetExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up CacheInterface.Get return parameters for the expectation previously defined by the When method
func (e *CacheInterfaceMockGetExpectation) Then(up1 *models.User, err error) *CacheInterfaceMock {
	e.results = &CacheInterfaceMockGetResults{up1, err}
	return e.mock
}

// Times sets number of times CacheInterface.Get should be invoked
func (mmGet *mCacheInterfaceMockGet) Times(n uint64) *mCacheInterfaceMockGet {
	if n == 0 {
		mmGet.mock.t.Fatalf("Times of CacheInterfaceMock.Get mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGet.expectedInvocations, n)
	mmGet.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGet
}

func (mmGet *mCacheInterfaceMockGet) invocationsDone() bool {
	if len(mmGet.expectations) == 0 && mmGet.defaultExpectation == nil && mmGet.mock.funcGet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGet.mock.afterGetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Get implements mm_repoInterface.CacheInterface
func (mmGet *CacheInterfaceMock) Get(ctx context.Context, params models.GetUserParams) (up1 *models.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	mmGet.t.Helper()

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, params)
	}

	mm_params := CacheInterfaceMockGetParams{ctx, params}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_want_ptrs := mmGet.GetMock.defaultExpectation.paramPtrs

		mm_got := CacheInterfaceMockGetParams{ctx, params}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGet.t.Errorf("CacheInterfaceMock.Get got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGet.GetMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.params != nil && !minimock.Equal(*mm_want_ptrs.params, mm_got.params) {
				mmGet.t.Errorf("CacheInterfaceMock.Get got unexpected parameter params, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGet.GetMock.defaultExpectation.expectationOrigins.originParams, *mm_want_ptrs.params, mm_got.params, minimock.Diff(*mm_want_ptrs.params, mm_got.params))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("CacheInterfaceMock.Get got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGet.GetMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the CacheInterfaceMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, params)
	}
	mmGet.t.Fatalf("Unexpected call to CacheInterfaceMock.Get. %v %v", ctx, params)
	return
}

// GetAfterCounter returns a count of finished CacheInterfaceMock.Get invocations
func (mmGet *CacheInterfaceMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of CacheInterfaceMock.Get invocations
func (mmGet *CacheInterfaceMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to CacheInterfaceMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mCacheInterfaceMockGet) Calls() []*CacheInterfaceMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*CacheInterfaceMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *CacheInterfaceMock) MinimockGetDone() bool {
	if m.GetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetMock.invocationsDone()
}

// MinimockGetInspect logs each unmet expectation
func (m *CacheInterfaceMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CacheInterfaceMock.Get at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetCounter := mm_atomic.LoadUint64(&m.afterGetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && afterGetCounter < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to CacheInterfaceMock.Get at\n%s", m.GetMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to CacheInterfaceMock.Get at\n%s with params: %#v", m.GetMock.defaultExpectation.expectationOrigins.origin, *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && afterGetCounter < 1 {
		m.t.Errorf("Expected call to CacheInterfaceMock.Get at\n%s", m.funcGetOrigin)
	}

	if !m.GetMock.invocationsDone() && afterGetCounter > 0 {
		m.t.Errorf("Expected %d calls to CacheInterfaceMock.Get at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetMock.expectedInvocations), m.GetMock.expectedInvocationsOrigin, afterGetCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CacheInterfaceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockGetInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CacheInterfaceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *CacheInterfaceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetDone()
}