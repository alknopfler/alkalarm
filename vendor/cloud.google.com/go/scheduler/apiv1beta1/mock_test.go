// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package scheduler

import (
	emptypb "github.com/golang/protobuf/ptypes/empty"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1beta1"
)

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockCloudSchedulerServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	schedulerpb.CloudSchedulerServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockCloudSchedulerServer) ListJobs(ctx context.Context, req *schedulerpb.ListJobsRequest) (*schedulerpb.ListJobsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*schedulerpb.ListJobsResponse), nil
}

func (s *mockCloudSchedulerServer) GetJob(ctx context.Context, req *schedulerpb.GetJobRequest) (*schedulerpb.Job, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*schedulerpb.Job), nil
}

func (s *mockCloudSchedulerServer) CreateJob(ctx context.Context, req *schedulerpb.CreateJobRequest) (*schedulerpb.Job, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*schedulerpb.Job), nil
}

func (s *mockCloudSchedulerServer) UpdateJob(ctx context.Context, req *schedulerpb.UpdateJobRequest) (*schedulerpb.Job, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*schedulerpb.Job), nil
}

func (s *mockCloudSchedulerServer) DeleteJob(ctx context.Context, req *schedulerpb.DeleteJobRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockCloudSchedulerServer) PauseJob(ctx context.Context, req *schedulerpb.PauseJobRequest) (*schedulerpb.Job, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*schedulerpb.Job), nil
}

func (s *mockCloudSchedulerServer) ResumeJob(ctx context.Context, req *schedulerpb.ResumeJobRequest) (*schedulerpb.Job, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*schedulerpb.Job), nil
}

func (s *mockCloudSchedulerServer) RunJob(ctx context.Context, req *schedulerpb.RunJobRequest) (*schedulerpb.Job, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*schedulerpb.Job), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockCloudScheduler mockCloudSchedulerServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	schedulerpb.RegisterCloudSchedulerServer(serv, &mockCloudScheduler)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestCloudSchedulerListJobs(t *testing.T) {
	var nextPageToken string = ""
	var jobsElement *schedulerpb.Job = &schedulerpb.Job{}
	var jobs = []*schedulerpb.Job{jobsElement}
	var expectedResponse = &schedulerpb.ListJobsResponse{
		NextPageToken: nextPageToken,
		Jobs:          jobs,
	}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &schedulerpb.ListJobsRequest{
		Parent: formattedParent,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListJobs(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Jobs[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCloudSchedulerListJobsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &schedulerpb.ListJobsRequest{
		Parent: formattedParent,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListJobs(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCloudSchedulerGetJob(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var schedule string = "schedule-697920873"
	var timeZone string = "timeZone36848094"
	var expectedResponse = &schedulerpb.Job{
		Name:        name2,
		Description: description,
		Schedule:    schedule,
		TimeZone:    timeZone,
	}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.GetJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetJob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCloudSchedulerGetJobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.GetJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetJob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCloudSchedulerCreateJob(t *testing.T) {
	var name string = "name3373707"
	var description string = "description-1724546052"
	var schedule string = "schedule-697920873"
	var timeZone string = "timeZone36848094"
	var expectedResponse = &schedulerpb.Job{
		Name:        name,
		Description: description,
		Schedule:    schedule,
		TimeZone:    timeZone,
	}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var job *schedulerpb.Job = &schedulerpb.Job{}
	var request = &schedulerpb.CreateJobRequest{
		Parent: formattedParent,
		Job:    job,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateJob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCloudSchedulerCreateJobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var job *schedulerpb.Job = &schedulerpb.Job{}
	var request = &schedulerpb.CreateJobRequest{
		Parent: formattedParent,
		Job:    job,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateJob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCloudSchedulerUpdateJob(t *testing.T) {
	var name string = "name3373707"
	var description string = "description-1724546052"
	var schedule string = "schedule-697920873"
	var timeZone string = "timeZone36848094"
	var expectedResponse = &schedulerpb.Job{
		Name:        name,
		Description: description,
		Schedule:    schedule,
		TimeZone:    timeZone,
	}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var job *schedulerpb.Job = &schedulerpb.Job{}
	var request = &schedulerpb.UpdateJobRequest{
		Job: job,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateJob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCloudSchedulerUpdateJobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var job *schedulerpb.Job = &schedulerpb.Job{}
	var request = &schedulerpb.UpdateJobRequest{
		Job: job,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateJob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCloudSchedulerDeleteJob(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.DeleteJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteJob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestCloudSchedulerDeleteJobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.DeleteJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteJob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestCloudSchedulerPauseJob(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var schedule string = "schedule-697920873"
	var timeZone string = "timeZone36848094"
	var expectedResponse = &schedulerpb.Job{
		Name:        name2,
		Description: description,
		Schedule:    schedule,
		TimeZone:    timeZone,
	}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.PauseJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.PauseJob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCloudSchedulerPauseJobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.PauseJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.PauseJob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCloudSchedulerResumeJob(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var schedule string = "schedule-697920873"
	var timeZone string = "timeZone36848094"
	var expectedResponse = &schedulerpb.Job{
		Name:        name2,
		Description: description,
		Schedule:    schedule,
		TimeZone:    timeZone,
	}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.ResumeJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ResumeJob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCloudSchedulerResumeJobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.ResumeJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ResumeJob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCloudSchedulerRunJob(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var schedule string = "schedule-697920873"
	var timeZone string = "timeZone36848094"
	var expectedResponse = &schedulerpb.Job{
		Name:        name2,
		Description: description,
		Schedule:    schedule,
		TimeZone:    timeZone,
	}

	mockCloudScheduler.err = nil
	mockCloudScheduler.reqs = nil

	mockCloudScheduler.resps = append(mockCloudScheduler.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.RunJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.RunJob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCloudScheduler.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCloudSchedulerRunJobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCloudScheduler.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", "[PROJECT]", "[LOCATION]", "[JOB]")
	var request = &schedulerpb.RunJobRequest{
		Name: formattedName,
	}

	c, err := NewCloudSchedulerClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.RunJob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
