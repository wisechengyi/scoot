// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package worker

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Status int64

const (
	Status_UNKNOWN  Status = 0
	Status_PENDING  Status = 1
	Status_RUNNING  Status = 2
	Status_COMPLETE Status = 3
	Status_FAILED   Status = 4
	Status_ABORTED  Status = 5
	Status_TIMEDOUT Status = 6
)

func (p Status) String() string {
	switch p {
	case Status_UNKNOWN:
		return "UNKNOWN"
	case Status_PENDING:
		return "PENDING"
	case Status_RUNNING:
		return "RUNNING"
	case Status_COMPLETE:
		return "COMPLETE"
	case Status_FAILED:
		return "FAILED"
	case Status_ABORTED:
		return "ABORTED"
	case Status_TIMEDOUT:
		return "TIMEDOUT"
	}
	return "<UNSET>"
}

func StatusFromString(s string) (Status, error) {
	switch s {
	case "UNKNOWN":
		return Status_UNKNOWN, nil
	case "PENDING":
		return Status_PENDING, nil
	case "RUNNING":
		return Status_RUNNING, nil
	case "COMPLETE":
		return Status_COMPLETE, nil
	case "FAILED":
		return Status_FAILED, nil
	case "ABORTED":
		return Status_ABORTED, nil
	case "TIMEDOUT":
		return Status_TIMEDOUT, nil
	}
	return Status(0), fmt.Errorf("not a valid Status string")
}

func StatusPtr(v Status) *Status { return &v }

func (p Status) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *Status) UnmarshalText(text []byte) error {
	q, err := StatusFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//   - Status
//   - RunId
//   - OutUri
//   - ErrUri
//   - Error
//   - ExitCode
//   - SnapshotId
//   - JobId
//   - TaskId
//   - Tag
type RunStatus struct {
	Status     Status  `thrift:"status,1,required" json:"status"`
	RunId      string  `thrift:"runId,2,required" json:"runId"`
	OutUri     *string `thrift:"outUri,3" json:"outUri,omitempty"`
	ErrUri     *string `thrift:"errUri,4" json:"errUri,omitempty"`
	Error      *string `thrift:"error,5" json:"error,omitempty"`
	ExitCode   *int32  `thrift:"exitCode,6" json:"exitCode,omitempty"`
	SnapshotId *string `thrift:"snapshotId,7" json:"snapshotId,omitempty"`
	JobId      *string `thrift:"jobId,8" json:"jobId,omitempty"`
	TaskId     *string `thrift:"taskId,9" json:"taskId,omitempty"`
	Tag        *string `thrift:"tag,10" json:"tag,omitempty"`
}

func NewRunStatus() *RunStatus {
	return &RunStatus{}
}

func (p *RunStatus) GetStatus() Status {
	return p.Status
}

func (p *RunStatus) GetRunId() string {
	return p.RunId
}

var RunStatus_OutUri_DEFAULT string

func (p *RunStatus) GetOutUri() string {
	if !p.IsSetOutUri() {
		return RunStatus_OutUri_DEFAULT
	}
	return *p.OutUri
}

var RunStatus_ErrUri_DEFAULT string

func (p *RunStatus) GetErrUri() string {
	if !p.IsSetErrUri() {
		return RunStatus_ErrUri_DEFAULT
	}
	return *p.ErrUri
}

var RunStatus_Error_DEFAULT string

func (p *RunStatus) GetError() string {
	if !p.IsSetError() {
		return RunStatus_Error_DEFAULT
	}
	return *p.Error
}

var RunStatus_ExitCode_DEFAULT int32

func (p *RunStatus) GetExitCode() int32 {
	if !p.IsSetExitCode() {
		return RunStatus_ExitCode_DEFAULT
	}
	return *p.ExitCode
}

var RunStatus_SnapshotId_DEFAULT string

func (p *RunStatus) GetSnapshotId() string {
	if !p.IsSetSnapshotId() {
		return RunStatus_SnapshotId_DEFAULT
	}
	return *p.SnapshotId
}

var RunStatus_JobId_DEFAULT string

func (p *RunStatus) GetJobId() string {
	if !p.IsSetJobId() {
		return RunStatus_JobId_DEFAULT
	}
	return *p.JobId
}

var RunStatus_TaskId_DEFAULT string

func (p *RunStatus) GetTaskId() string {
	if !p.IsSetTaskId() {
		return RunStatus_TaskId_DEFAULT
	}
	return *p.TaskId
}

var RunStatus_Tag_DEFAULT string

func (p *RunStatus) GetTag() string {
	if !p.IsSetTag() {
		return RunStatus_Tag_DEFAULT
	}
	return *p.Tag
}
func (p *RunStatus) IsSetOutUri() bool {
	return p.OutUri != nil
}

func (p *RunStatus) IsSetErrUri() bool {
	return p.ErrUri != nil
}

func (p *RunStatus) IsSetError() bool {
	return p.Error != nil
}

func (p *RunStatus) IsSetExitCode() bool {
	return p.ExitCode != nil
}

func (p *RunStatus) IsSetSnapshotId() bool {
	return p.SnapshotId != nil
}

func (p *RunStatus) IsSetJobId() bool {
	return p.JobId != nil
}

func (p *RunStatus) IsSetTaskId() bool {
	return p.TaskId != nil
}

func (p *RunStatus) IsSetTag() bool {
	return p.Tag != nil
}

func (p *RunStatus) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetStatus bool = false
	var issetRunId bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetStatus = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetRunId = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.readField9(iprot); err != nil {
				return err
			}
		case 10:
			if err := p.readField10(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetStatus {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Status is not set"))
	}
	if !issetRunId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RunId is not set"))
	}
	return nil
}

func (p *RunStatus) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := Status(v)
		p.Status = temp
	}
	return nil
}

func (p *RunStatus) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.RunId = v
	}
	return nil
}

func (p *RunStatus) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.OutUri = &v
	}
	return nil
}

func (p *RunStatus) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.ErrUri = &v
	}
	return nil
}

func (p *RunStatus) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Error = &v
	}
	return nil
}

func (p *RunStatus) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.ExitCode = &v
	}
	return nil
}

func (p *RunStatus) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.SnapshotId = &v
	}
	return nil
}

func (p *RunStatus) readField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.JobId = &v
	}
	return nil
}

func (p *RunStatus) readField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.TaskId = &v
	}
	return nil
}

func (p *RunStatus) readField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.Tag = &v
	}
	return nil
}

func (p *RunStatus) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RunStatus"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := p.writeField10(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunStatus) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:status: ", p), err)
	}
	return err
}

func (p *RunStatus) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("runId", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:runId: ", p), err)
	}
	if err := oprot.WriteString(string(p.RunId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.runId (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:runId: ", p), err)
	}
	return err
}

func (p *RunStatus) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetOutUri() {
		if err := oprot.WriteFieldBegin("outUri", thrift.STRING, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:outUri: ", p), err)
		}
		if err := oprot.WriteString(string(*p.OutUri)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.outUri (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:outUri: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrUri() {
		if err := oprot.WriteFieldBegin("errUri", thrift.STRING, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:errUri: ", p), err)
		}
		if err := oprot.WriteString(string(*p.ErrUri)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.errUri (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:errUri: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetError() {
		if err := oprot.WriteFieldBegin("error", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:error: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Error)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.error (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:error: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetExitCode() {
		if err := oprot.WriteFieldBegin("exitCode", thrift.I32, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:exitCode: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.ExitCode)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.exitCode (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:exitCode: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetSnapshotId() {
		if err := oprot.WriteFieldBegin("snapshotId", thrift.STRING, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:snapshotId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.SnapshotId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.snapshotId (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:snapshotId: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetJobId() {
		if err := oprot.WriteFieldBegin("jobId", thrift.STRING, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:jobId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.JobId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.jobId (8) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:jobId: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) writeField9(oprot thrift.TProtocol) (err error) {
	if p.IsSetTaskId() {
		if err := oprot.WriteFieldBegin("taskId", thrift.STRING, 9); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:taskId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.TaskId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.taskId (9) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 9:taskId: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) writeField10(oprot thrift.TProtocol) (err error) {
	if p.IsSetTag() {
		if err := oprot.WriteFieldBegin("tag", thrift.STRING, 10); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:tag: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Tag)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.tag (10) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 10:tag: ", p), err)
		}
	}
	return err
}

func (p *RunStatus) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunStatus(%+v)", *p)
}

// Attributes:
//   - Runs
//   - Initialized
//   - IsHealthy
//   - Error
type WorkerStatus struct {
	Runs        []*RunStatus `thrift:"runs,1,required" json:"runs"`
	Initialized bool         `thrift:"initialized,2,required" json:"initialized"`
	IsHealthy   bool         `thrift:"isHealthy,3,required" json:"isHealthy"`
	Error       string       `thrift:"error,4,required" json:"error"`
}

func NewWorkerStatus() *WorkerStatus {
	return &WorkerStatus{}
}

func (p *WorkerStatus) GetRuns() []*RunStatus {
	return p.Runs
}

func (p *WorkerStatus) GetInitialized() bool {
	return p.Initialized
}

func (p *WorkerStatus) GetIsHealthy() bool {
	return p.IsHealthy
}

func (p *WorkerStatus) GetError() string {
	return p.Error
}
func (p *WorkerStatus) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRuns bool = false
	var issetInitialized bool = false
	var issetIsHealthy bool = false
	var issetError bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRuns = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetInitialized = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetIsHealthy = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
			issetError = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRuns {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Runs is not set"))
	}
	if !issetInitialized {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Initialized is not set"))
	}
	if !issetIsHealthy {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field IsHealthy is not set"))
	}
	if !issetError {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Error is not set"))
	}
	return nil
}

func (p *WorkerStatus) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*RunStatus, 0, size)
	p.Runs = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &RunStatus{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Runs = append(p.Runs, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *WorkerStatus) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Initialized = v
	}
	return nil
}

func (p *WorkerStatus) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.IsHealthy = v
	}
	return nil
}

func (p *WorkerStatus) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Error = v
	}
	return nil
}

func (p *WorkerStatus) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("WorkerStatus"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerStatus) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("runs", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:runs: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Runs)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Runs {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:runs: ", p), err)
	}
	return err
}

func (p *WorkerStatus) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("initialized", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:initialized: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.Initialized)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.initialized (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:initialized: ", p), err)
	}
	return err
}

func (p *WorkerStatus) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("isHealthy", thrift.BOOL, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:isHealthy: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.IsHealthy)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.isHealthy (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:isHealthy: ", p), err)
	}
	return err
}

func (p *WorkerStatus) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("error", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:error: ", p), err)
	}
	if err := oprot.WriteString(string(p.Error)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.error (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:error: ", p), err)
	}
	return err
}

func (p *WorkerStatus) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerStatus(%+v)", *p)
}

// Attributes:
//   - Argv
//   - Env
//   - SnapshotId
//   - TimeoutMs
//   - JobId
//   - TaskId
//   - Tag
type RunCommand struct {
	Argv       []string          `thrift:"argv,1,required" json:"argv"`
	Env        map[string]string `thrift:"env,2" json:"env,omitempty"`
	SnapshotId *string           `thrift:"snapshotId,3" json:"snapshotId,omitempty"`
	TimeoutMs  *int32            `thrift:"timeoutMs,4" json:"timeoutMs,omitempty"`
	JobId      *string           `thrift:"jobId,5" json:"jobId,omitempty"`
	TaskId     *string           `thrift:"taskId,6" json:"taskId,omitempty"`
	Tag        *string           `thrift:"tag,7" json:"tag,omitempty"`
}

func NewRunCommand() *RunCommand {
	return &RunCommand{}
}

func (p *RunCommand) GetArgv() []string {
	return p.Argv
}

var RunCommand_Env_DEFAULT map[string]string

func (p *RunCommand) GetEnv() map[string]string {
	return p.Env
}

var RunCommand_SnapshotId_DEFAULT string

func (p *RunCommand) GetSnapshotId() string {
	if !p.IsSetSnapshotId() {
		return RunCommand_SnapshotId_DEFAULT
	}
	return *p.SnapshotId
}

var RunCommand_TimeoutMs_DEFAULT int32

func (p *RunCommand) GetTimeoutMs() int32 {
	if !p.IsSetTimeoutMs() {
		return RunCommand_TimeoutMs_DEFAULT
	}
	return *p.TimeoutMs
}

var RunCommand_JobId_DEFAULT string

func (p *RunCommand) GetJobId() string {
	if !p.IsSetJobId() {
		return RunCommand_JobId_DEFAULT
	}
	return *p.JobId
}

var RunCommand_TaskId_DEFAULT string

func (p *RunCommand) GetTaskId() string {
	if !p.IsSetTaskId() {
		return RunCommand_TaskId_DEFAULT
	}
	return *p.TaskId
}

var RunCommand_Tag_DEFAULT string

func (p *RunCommand) GetTag() string {
	if !p.IsSetTag() {
		return RunCommand_Tag_DEFAULT
	}
	return *p.Tag
}
func (p *RunCommand) IsSetEnv() bool {
	return p.Env != nil
}

func (p *RunCommand) IsSetSnapshotId() bool {
	return p.SnapshotId != nil
}

func (p *RunCommand) IsSetTimeoutMs() bool {
	return p.TimeoutMs != nil
}

func (p *RunCommand) IsSetJobId() bool {
	return p.JobId != nil
}

func (p *RunCommand) IsSetTaskId() bool {
	return p.TaskId != nil
}

func (p *RunCommand) IsSetTag() bool {
	return p.Tag != nil
}

func (p *RunCommand) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetArgv bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetArgv = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetArgv {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Argv is not set"))
	}
	return nil
}

func (p *RunCommand) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Argv = tSlice
	for i := 0; i < size; i++ {
		var _elem1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.Argv = append(p.Argv, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *RunCommand) readField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Env = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key2 = v
		}
		var _val3 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val3 = v
		}
		p.Env[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *RunCommand) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.SnapshotId = &v
	}
	return nil
}

func (p *RunCommand) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.TimeoutMs = &v
	}
	return nil
}

func (p *RunCommand) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.JobId = &v
	}
	return nil
}

func (p *RunCommand) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.TaskId = &v
	}
	return nil
}

func (p *RunCommand) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Tag = &v
	}
	return nil
}

func (p *RunCommand) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RunCommand"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunCommand) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("argv", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:argv: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Argv)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Argv {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:argv: ", p), err)
	}
	return err
}

func (p *RunCommand) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetEnv() {
		if err := oprot.WriteFieldBegin("env", thrift.MAP, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:env: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Env)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Env {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:env: ", p), err)
		}
	}
	return err
}

func (p *RunCommand) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetSnapshotId() {
		if err := oprot.WriteFieldBegin("snapshotId", thrift.STRING, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:snapshotId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.SnapshotId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.snapshotId (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:snapshotId: ", p), err)
		}
	}
	return err
}

func (p *RunCommand) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimeoutMs() {
		if err := oprot.WriteFieldBegin("timeoutMs", thrift.I32, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:timeoutMs: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.TimeoutMs)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timeoutMs (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:timeoutMs: ", p), err)
		}
	}
	return err
}

func (p *RunCommand) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetJobId() {
		if err := oprot.WriteFieldBegin("jobId", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:jobId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.JobId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.jobId (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:jobId: ", p), err)
		}
	}
	return err
}

func (p *RunCommand) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetTaskId() {
		if err := oprot.WriteFieldBegin("taskId", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:taskId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.TaskId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.taskId (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:taskId: ", p), err)
		}
	}
	return err
}

func (p *RunCommand) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetTag() {
		if err := oprot.WriteFieldBegin("tag", thrift.STRING, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:tag: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Tag)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.tag (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:tag: ", p), err)
		}
	}
	return err
}

func (p *RunCommand) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunCommand(%+v)", *p)
}
