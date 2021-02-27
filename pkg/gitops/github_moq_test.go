// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package gitops

import (
	"context"
	"sync"
)

// Ensure, that pullRequestOpenerMock does implement pullRequestOpener.
// If this is not the case, regenerate this file with moq.
var _ pullRequestOpener = &pullRequestOpenerMock{}

// pullRequestOpenerMock is a mock implementation of pullRequestOpener.
//
//     func TestSomethingThatUsespullRequestOpener(t *testing.T) {
//
//         // make and configure a mocked pullRequestOpener
//         mockedpullRequestOpener := &pullRequestOpenerMock{
//             OpenPullRequestFunc: func(in1 context.Context, in2 openPullRequestParams) (string, error) {
// 	               panic("mock out the OpenPullRequest method")
//             },
//         }
//
//         // use mockedpullRequestOpener in code that requires pullRequestOpener
//         // and then make assertions.
//
//     }
type pullRequestOpenerMock struct {
	// OpenPullRequestFunc mocks the OpenPullRequest method.
	OpenPullRequestFunc func(in1 context.Context, in2 openPullRequestParams) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// OpenPullRequest holds details about calls to the OpenPullRequest method.
		OpenPullRequest []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 openPullRequestParams
		}
	}
	lockOpenPullRequest sync.RWMutex
}

// OpenPullRequest calls OpenPullRequestFunc.
func (mock *pullRequestOpenerMock) OpenPullRequest(in1 context.Context, in2 openPullRequestParams) (string, error) {
	if mock.OpenPullRequestFunc == nil {
		panic("pullRequestOpenerMock.OpenPullRequestFunc: method is nil but pullRequestOpener.OpenPullRequest was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 openPullRequestParams
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockOpenPullRequest.Lock()
	mock.calls.OpenPullRequest = append(mock.calls.OpenPullRequest, callInfo)
	mock.lockOpenPullRequest.Unlock()
	return mock.OpenPullRequestFunc(in1, in2)
}

// OpenPullRequestCalls gets all the calls that were made to OpenPullRequest.
// Check the length with:
//     len(mockedpullRequestOpener.OpenPullRequestCalls())
func (mock *pullRequestOpenerMock) OpenPullRequestCalls() []struct {
	In1 context.Context
	In2 openPullRequestParams
} {
	var calls []struct {
		In1 context.Context
		In2 openPullRequestParams
	}
	mock.lockOpenPullRequest.RLock()
	calls = mock.calls.OpenPullRequest
	mock.lockOpenPullRequest.RUnlock()
	return calls
}
