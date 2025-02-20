// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Dnlbb/auth/internal/producer.Producer -o producer_minimock.go -n ProducerMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/IBM/sarama"
	"github.com/gojuno/minimock/v3"
)

// ProducerMock implements mm_producer.Producer
type ProducerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcClose          func() (err error)
	funcCloseOrigin    string
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mProducerMockClose

	funcSendMessage          func(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
	funcSendMessageOrigin    string
	inspectFuncSendMessage   func(msg *sarama.ProducerMessage)
	afterSendMessageCounter  uint64
	beforeSendMessageCounter uint64
	SendMessageMock          mProducerMockSendMessage
}

// NewProducerMock returns a mock for mm_producer.Producer
func NewProducerMock(t minimock.Tester) *ProducerMock {
	m := &ProducerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mProducerMockClose{mock: m}

	m.SendMessageMock = mProducerMockSendMessage{mock: m}
	m.SendMessageMock.callArgs = []*ProducerMockSendMessageParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mProducerMockClose struct {
	optional           bool
	mock               *ProducerMock
	defaultExpectation *ProducerMockCloseExpectation
	expectations       []*ProducerMockCloseExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ProducerMockCloseExpectation specifies expectation struct of the Producer.Close
type ProducerMockCloseExpectation struct {
	mock *ProducerMock

	results      *ProducerMockCloseResults
	returnOrigin string
	Counter      uint64
}

// ProducerMockCloseResults contains results of the Producer.Close
type ProducerMockCloseResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmClose *mProducerMockClose) Optional() *mProducerMockClose {
	mmClose.optional = true
	return mmClose
}

// Expect sets up expected params for Producer.Close
func (mmClose *mProducerMockClose) Expect() *mProducerMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("ProducerMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &ProducerMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the Producer.Close
func (mmClose *mProducerMockClose) Inspect(f func()) *mProducerMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for ProducerMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by Producer.Close
func (mmClose *mProducerMockClose) Return(err error) *ProducerMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("ProducerMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &ProducerMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &ProducerMockCloseResults{err}
	mmClose.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Set uses given function f to mock the Producer.Close method
func (mmClose *mProducerMockClose) Set(f func() (err error)) *ProducerMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the Producer.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the Producer.Close method")
	}

	mmClose.mock.funcClose = f
	mmClose.mock.funcCloseOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Times sets number of times Producer.Close should be invoked
func (mmClose *mProducerMockClose) Times(n uint64) *mProducerMockClose {
	if n == 0 {
		mmClose.mock.t.Fatalf("Times of ProducerMock.Close mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmClose.expectedInvocations, n)
	mmClose.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmClose
}

func (mmClose *mProducerMockClose) invocationsDone() bool {
	if len(mmClose.expectations) == 0 && mmClose.defaultExpectation == nil && mmClose.mock.funcClose == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmClose.mock.afterCloseCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmClose.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Close implements mm_producer.Producer
func (mmClose *ProducerMock) Close() (err error) {
	mm_atomic.AddUint64(&mmClose.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&mmClose.afterCloseCounter, 1)

	mmClose.t.Helper()

	if mmClose.inspectFuncClose != nil {
		mmClose.inspectFuncClose()
	}

	if mmClose.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClose.CloseMock.defaultExpectation.Counter, 1)

		mm_results := mmClose.CloseMock.defaultExpectation.results
		if mm_results == nil {
			mmClose.t.Fatal("No results are set for the ProducerMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to ProducerMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished ProducerMock.Close invocations
func (mmClose *ProducerMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of ProducerMock.Close invocations
func (mmClose *ProducerMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *ProducerMock) MinimockCloseDone() bool {
	if m.CloseMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CloseMock.invocationsDone()
}

// MinimockCloseInspect logs each unmet expectation
func (m *ProducerMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ProducerMock.Close")
		}
	}

	afterCloseCounter := mm_atomic.LoadUint64(&m.afterCloseCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to ProducerMock.Close at\n%s", m.CloseMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to ProducerMock.Close at\n%s", m.funcCloseOrigin)
	}

	if !m.CloseMock.invocationsDone() && afterCloseCounter > 0 {
		m.t.Errorf("Expected %d calls to ProducerMock.Close at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CloseMock.expectedInvocations), m.CloseMock.expectedInvocationsOrigin, afterCloseCounter)
	}
}

type mProducerMockSendMessage struct {
	optional           bool
	mock               *ProducerMock
	defaultExpectation *ProducerMockSendMessageExpectation
	expectations       []*ProducerMockSendMessageExpectation

	callArgs []*ProducerMockSendMessageParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ProducerMockSendMessageExpectation specifies expectation struct of the Producer.SendMessage
type ProducerMockSendMessageExpectation struct {
	mock               *ProducerMock
	params             *ProducerMockSendMessageParams
	paramPtrs          *ProducerMockSendMessageParamPtrs
	expectationOrigins ProducerMockSendMessageExpectationOrigins
	results            *ProducerMockSendMessageResults
	returnOrigin       string
	Counter            uint64
}

// ProducerMockSendMessageParams contains parameters of the Producer.SendMessage
type ProducerMockSendMessageParams struct {
	msg *sarama.ProducerMessage
}

// ProducerMockSendMessageParamPtrs contains pointers to parameters of the Producer.SendMessage
type ProducerMockSendMessageParamPtrs struct {
	msg **sarama.ProducerMessage
}

// ProducerMockSendMessageResults contains results of the Producer.SendMessage
type ProducerMockSendMessageResults struct {
	partition int32
	offset    int64
	err       error
}

// ProducerMockSendMessageOrigins contains origins of expectations of the Producer.SendMessage
type ProducerMockSendMessageExpectationOrigins struct {
	origin    string
	originMsg string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSendMessage *mProducerMockSendMessage) Optional() *mProducerMockSendMessage {
	mmSendMessage.optional = true
	return mmSendMessage
}

// Expect sets up expected params for Producer.SendMessage
func (mmSendMessage *mProducerMockSendMessage) Expect(msg *sarama.ProducerMessage) *mProducerMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ProducerMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &ProducerMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.paramPtrs != nil {
		mmSendMessage.mock.t.Fatalf("ProducerMock.SendMessage mock is already set by ExpectParams functions")
	}

	mmSendMessage.defaultExpectation.params = &ProducerMockSendMessageParams{msg}
	mmSendMessage.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmSendMessage.expectations {
		if minimock.Equal(e.params, mmSendMessage.defaultExpectation.params) {
			mmSendMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMessage.defaultExpectation.params)
		}
	}

	return mmSendMessage
}

// ExpectMsgParam1 sets up expected param msg for Producer.SendMessage
func (mmSendMessage *mProducerMockSendMessage) ExpectMsgParam1(msg *sarama.ProducerMessage) *mProducerMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ProducerMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &ProducerMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.params != nil {
		mmSendMessage.mock.t.Fatalf("ProducerMock.SendMessage mock is already set by Expect")
	}

	if mmSendMessage.defaultExpectation.paramPtrs == nil {
		mmSendMessage.defaultExpectation.paramPtrs = &ProducerMockSendMessageParamPtrs{}
	}
	mmSendMessage.defaultExpectation.paramPtrs.msg = &msg
	mmSendMessage.defaultExpectation.expectationOrigins.originMsg = minimock.CallerInfo(1)

	return mmSendMessage
}

// Inspect accepts an inspector function that has same arguments as the Producer.SendMessage
func (mmSendMessage *mProducerMockSendMessage) Inspect(f func(msg *sarama.ProducerMessage)) *mProducerMockSendMessage {
	if mmSendMessage.mock.inspectFuncSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("Inspect function is already set for ProducerMock.SendMessage")
	}

	mmSendMessage.mock.inspectFuncSendMessage = f

	return mmSendMessage
}

// Return sets up results that will be returned by Producer.SendMessage
func (mmSendMessage *mProducerMockSendMessage) Return(partition int32, offset int64, err error) *ProducerMock {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ProducerMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &ProducerMockSendMessageExpectation{mock: mmSendMessage.mock}
	}
	mmSendMessage.defaultExpectation.results = &ProducerMockSendMessageResults{partition, offset, err}
	mmSendMessage.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSendMessage.mock
}

// Set uses given function f to mock the Producer.SendMessage method
func (mmSendMessage *mProducerMockSendMessage) Set(f func(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)) *ProducerMock {
	if mmSendMessage.defaultExpectation != nil {
		mmSendMessage.mock.t.Fatalf("Default expectation is already set for the Producer.SendMessage method")
	}

	if len(mmSendMessage.expectations) > 0 {
		mmSendMessage.mock.t.Fatalf("Some expectations are already set for the Producer.SendMessage method")
	}

	mmSendMessage.mock.funcSendMessage = f
	mmSendMessage.mock.funcSendMessageOrigin = minimock.CallerInfo(1)
	return mmSendMessage.mock
}

// When sets expectation for the Producer.SendMessage which will trigger the result defined by the following
// Then helper
func (mmSendMessage *mProducerMockSendMessage) When(msg *sarama.ProducerMessage) *ProducerMockSendMessageExpectation {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ProducerMock.SendMessage mock is already set by Set")
	}

	expectation := &ProducerMockSendMessageExpectation{
		mock:               mmSendMessage.mock,
		params:             &ProducerMockSendMessageParams{msg},
		expectationOrigins: ProducerMockSendMessageExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmSendMessage.expectations = append(mmSendMessage.expectations, expectation)
	return expectation
}

// Then sets up Producer.SendMessage return parameters for the expectation previously defined by the When method
func (e *ProducerMockSendMessageExpectation) Then(partition int32, offset int64, err error) *ProducerMock {
	e.results = &ProducerMockSendMessageResults{partition, offset, err}
	return e.mock
}

// Times sets number of times Producer.SendMessage should be invoked
func (mmSendMessage *mProducerMockSendMessage) Times(n uint64) *mProducerMockSendMessage {
	if n == 0 {
		mmSendMessage.mock.t.Fatalf("Times of ProducerMock.SendMessage mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSendMessage.expectedInvocations, n)
	mmSendMessage.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSendMessage
}

func (mmSendMessage *mProducerMockSendMessage) invocationsDone() bool {
	if len(mmSendMessage.expectations) == 0 && mmSendMessage.defaultExpectation == nil && mmSendMessage.mock.funcSendMessage == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSendMessage.mock.afterSendMessageCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSendMessage.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SendMessage implements mm_producer.Producer
func (mmSendMessage *ProducerMock) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	mm_atomic.AddUint64(&mmSendMessage.beforeSendMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMessage.afterSendMessageCounter, 1)

	mmSendMessage.t.Helper()

	if mmSendMessage.inspectFuncSendMessage != nil {
		mmSendMessage.inspectFuncSendMessage(msg)
	}

	mm_params := ProducerMockSendMessageParams{msg}

	// Record call args
	mmSendMessage.SendMessageMock.mutex.Lock()
	mmSendMessage.SendMessageMock.callArgs = append(mmSendMessage.SendMessageMock.callArgs, &mm_params)
	mmSendMessage.SendMessageMock.mutex.Unlock()

	for _, e := range mmSendMessage.SendMessageMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.partition, e.results.offset, e.results.err
		}
	}

	if mmSendMessage.SendMessageMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendMessage.SendMessageMock.defaultExpectation.Counter, 1)
		mm_want := mmSendMessage.SendMessageMock.defaultExpectation.params
		mm_want_ptrs := mmSendMessage.SendMessageMock.defaultExpectation.paramPtrs

		mm_got := ProducerMockSendMessageParams{msg}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.msg != nil && !minimock.Equal(*mm_want_ptrs.msg, mm_got.msg) {
				mmSendMessage.t.Errorf("ProducerMock.SendMessage got unexpected parameter msg, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSendMessage.SendMessageMock.defaultExpectation.expectationOrigins.originMsg, *mm_want_ptrs.msg, mm_got.msg, minimock.Diff(*mm_want_ptrs.msg, mm_got.msg))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMessage.t.Errorf("ProducerMock.SendMessage got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmSendMessage.SendMessageMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendMessage.SendMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmSendMessage.t.Fatal("No results are set for the ProducerMock.SendMessage")
		}
		return (*mm_results).partition, (*mm_results).offset, (*mm_results).err
	}
	if mmSendMessage.funcSendMessage != nil {
		return mmSendMessage.funcSendMessage(msg)
	}
	mmSendMessage.t.Fatalf("Unexpected call to ProducerMock.SendMessage. %v", msg)
	return
}

// SendMessageAfterCounter returns a count of finished ProducerMock.SendMessage invocations
func (mmSendMessage *ProducerMock) SendMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.afterSendMessageCounter)
}

// SendMessageBeforeCounter returns a count of ProducerMock.SendMessage invocations
func (mmSendMessage *ProducerMock) SendMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.beforeSendMessageCounter)
}

// Calls returns a list of arguments used in each call to ProducerMock.SendMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMessage *mProducerMockSendMessage) Calls() []*ProducerMockSendMessageParams {
	mmSendMessage.mutex.RLock()

	argCopy := make([]*ProducerMockSendMessageParams, len(mmSendMessage.callArgs))
	copy(argCopy, mmSendMessage.callArgs)

	mmSendMessage.mutex.RUnlock()

	return argCopy
}

// MinimockSendMessageDone returns true if the count of the SendMessage invocations corresponds
// the number of defined expectations
func (m *ProducerMock) MinimockSendMessageDone() bool {
	if m.SendMessageMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SendMessageMock.invocationsDone()
}

// MinimockSendMessageInspect logs each unmet expectation
func (m *ProducerMock) MinimockSendMessageInspect() {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ProducerMock.SendMessage at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterSendMessageCounter := mm_atomic.LoadUint64(&m.afterSendMessageCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && afterSendMessageCounter < 1 {
		if m.SendMessageMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to ProducerMock.SendMessage at\n%s", m.SendMessageMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to ProducerMock.SendMessage at\n%s with params: %#v", m.SendMessageMock.defaultExpectation.expectationOrigins.origin, *m.SendMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && afterSendMessageCounter < 1 {
		m.t.Errorf("Expected call to ProducerMock.SendMessage at\n%s", m.funcSendMessageOrigin)
	}

	if !m.SendMessageMock.invocationsDone() && afterSendMessageCounter > 0 {
		m.t.Errorf("Expected %d calls to ProducerMock.SendMessage at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SendMessageMock.expectedInvocations), m.SendMessageMock.expectedInvocationsOrigin, afterSendMessageCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ProducerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCloseInspect()

			m.MinimockSendMessageInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ProducerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ProducerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockSendMessageDone()
}
