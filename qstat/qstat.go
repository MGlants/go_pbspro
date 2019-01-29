package qstat

/*
#cgo CFLAGS: -g
#cgo LDFLAGS: -L/opt/pbspro/lib -lpbs
#include <stdlib.h>
#include "/opt/pbspro/include/pbs_error.h"
#include "/opt/pbspro/include/pbs_ifl.h"
*/
import "C"
import (
	"unsafe"

	"github.com/juju/errors"

	"github.com/taylor840326/go_pbspro/utils"
)

type (
	// qstat gather server state information.
	QstatServerInfo struct {
		ServerState             string `json:"server_state" db:"server_state"`
		ServerHost              string `json:"server_host" db:"server_host"`
		ServerScheduling        string `json:"server_scheduling" db:"server_scheduling"`
		TotalJobs               int64  `json:"total_jobs" db:"total_jobs"`
		StateCountTransit       int64  `json:"state_count_transit" db:"state_count_transit"`
		StateCountQueued        int64  `json:"state_count_queued" db:"state_count_queued"`
		StateCountHeld          int64  `json:"state_count_held" db:"state_count_held"`
		StateCountWaiting       int64  `json:"state_count_waiting" db:"state_count_waiting"`
		StateCountRunning       int64  `json:"state_count_running" db:"state_count_running"`
		StateCountExiting       int64  `json:"state_count_exiting" db:"state_count_exiting"`
		StateCountBegun         int64  `json:"state_count_begun" db:"state_count_begun"`
		DefaultQueue            string `json:"default_queue" db:"default_queue"`
		LogEvents               int64  `json:"log_events" db:"log_events"`
		MailFrom                string `json:"mail_from" db:"mail_from"`
		QueryOtherJobs          string `json:"query_other_jobs" db:"query_other_jobs"`
		ResourcesDefaultNcpus   int64  `json:"resources_default_ncpus" db:"resources_default_ncpus"`
		DefaultChunkNcpus       int64  `json:"default_chunk_ncpus" db:"default_chunk_ncpus"`
		ResourcesAssignedNcpus  int64  `json:"resources_assigned_ncpus" db:"resources_assigned_ncpus"`
		ResourcesAssignedNodect int64  `json:"resources_assigned_nodect" db:"resources_assigned_nodect"`
		SchedulerIteration      int64  `json:"scheduler_iteration" db:" scheduler_iteration"`
		Flicenses               int64  `json:"flicenses" db:"flicenses"`
		ResvEnable              string `json:"resv_enable" db:"resv_enable"`
		NodeFailRequeue         int64  `json:"node_fail_requeue" db:"node_fail_requeue"`
		MaxArraySize            int64  `json:"max_array_size" db:"max_array_size"`
		PBSLicenseMin           int64  `json:"pbs_license_min" db:"pbs_license_min"`
		PBSLicenseMax           int64  `json:"pbs_license_max" db:"pbs_license_max"`
		PBSLicenseLingerTime    int64  `json:"pbs_license_linger_time" db:"pbs_license_linger_time"`
		LicenseCountAvailGlobal int64  `json:"license_count_avail_global" db:"license_count_avail_global"`
		LicenseCountAvailLocal  int64  `json:"license_count_avail_local" db:"license_count_avail_local"`
		LicenseCountUsed        int64  `json:"license_count_used" db:"license_count_used"`
		LicenseCountHighUse     int64  `json:"license_count_high_use" db:"license_count_high_use"`
		PBSVersion              string `json:"pbs_version" db:"pbs_version"`
		EligibleTimeEnable      string `json:"eligible_time_enable" db:"eligible_time_enable"`
		JobHistoryEnable        string `json:"job_history_enable" db:"job_history_enable"`
		JobHistoryDuration      string `json:"job_history_duration" db:"job_history_duration"`
		MaxConcurrentProvision  int64  `json:"max_concurrent_provision" db:"max_concurrent_provision"`
	}

	// qstat gather queue information.
	QstatQueueInfo struct {
		QueueType               string `json:"queue_type" db:"queue_type"`
		TotalJobs               int64  `json:"total_jobs" db:"total_jobs"`
		StateCountTransit       int64  `json:"state_count_transit" db:"state_count_transit"`
		StateCountQueued        int64  `json:"state_count_queued" db:"state_count_queued"`
		StateCountHeld          int64  `json:"state_count_held" db:"state_count_held"`
		StateCountWaiting       int64  `json:"state_count_waiting" db:"state_count_waiting"`
		StateCountRunning       int64  `json:"state_count_running" db:"state_count_running"`
		StateCountExiting       int64  `json:"state_count_exiting" db:"state_count_exiting"`
		StateCountBegun         int64  `json:"state_count_begun" db:"state_count_begun"`
		ResourcesAssignedNcpus  int64  `json:"resources_assigned_ncpus" db:"resources_assigned_ncpus"`
		ResourcesAssignedNodect int64  `json:"resources_assigned_nodect" db:"resources_assigned_nodect"`
		Enable                  string `json:"enable" db:"enable"`
		Started                 string `json:"started" db:"started"`
	}

	//qstat gather node information.
	QstatNodeInfo struct {
		Mom                                string           `json:"mom" db:"mom"`
		Ntype                              string           `json:"ntype" db:"ntype"`
		State                              string           `json:'state" db:"state"`
		Pcpus                              int64            `json:"pcpus" db:"pcpus"`
		Jobs                               map[string]int64 `json:"jobs" db:"jobs"`
		ResourcesAvailableArch             string           `json:"resources_available_arch" db:"resources_available_arch"`
		ResourcesAvailableHost             string           `json:"resources_available_host" db:"resources_available_host"`
		ResourcesAvailableMem              string           `json:"resources_available_mem" db:"resources_available_mem"`
		ResourcesAvailableNcpus            int64            `json:"resources_available_ncpus" db:"resources_available_ncpus"`
		ResourcesAvailableApplications     string           `json:"resources_available_pas_applications_enabled" db:"resources_available_pas_applications_enabled"`
		ResourcesAvailablePlatform         string           `json:"resources_available_platform" db:"resources_available_platform"`
		ResourcesAvailableSoftware         string           `json:"resources_availabled_software" db:"resources_available_software"`
		ResourcesAvailableVnodes           string           `json:"resources_available_vnodes" db:"resources_available_vnodes"`
		ResourcesAssignedAcceleratorMemory string           `json:"resources_assigned_accelerator_memory" db:"resources_assigned_accelerator_memory"`
		ResourcesAssignedHbmem             string           `json:"resources_assigned_hbmem" db:"resources_assigned_hbmem"`
		ResourcesAssignedMem               string           `json:"resources_assigned_mem" db:"resources_assigned_mem"`
		ResourcesAssignedNaccelerators     int64            `json:"resources_assigned_naccelerators" db:"resources_assigned_naccelerators"`
		ResourcesAssignedNcpus             int64            `json:"resources_assigned_ncpus" db:"resources_assigned_ncpus"`
		ResourcesAssignedVmem              string           `json:"resources_assigned_vmem" db:"resources_assigned_vmem"`
		ResvEnable                         string           `json:"resv_enable" db:"resv_enable"`
		Sharing                            string           `json:"sharing" db:"sharing"`
		LastStateChangeTime                int64            `json:"last_state_change_time" db:"last_state_change_time"`
		LastUsedTime                       int64            `json:"last_used_time" db:"last_used_time"`
	}

	//定义PBS结构体
	Qstat struct {
		Server        string         `json:"server"`
		Handle        int            `json:"handle"`
		DefaultServer string         `json:"default_server"`
		IsClosed      bool           `json:"is_closed"`
		Attribs       []utils.Attrib `json:"attribs"`
		Extend        string         `json:"extend"`
		Id            string         `json:"id"`
	}
)

//新建一个Qstat实例
func NewQstat(server string) (qs *Qstat, err error) {
	qstat := new(Qstat)

	qstat.Server = server
	qstat.Handle = 0
	qstat.DefaultServer = ""
	qstat.IsClosed = false
	qstat.Attribs = nil
	qstat.Extend = ""
	qstat.Id = ""

	return qstat, nil
}

//设定服务名称
func (qs *Qstat) SetServerName(server string) {
	qs.Server = server
}

//设定handle号，>= 0
func (qs *Qstat) SetHandle(handle int) {
	qs.Handle = handle
}

//设定属性列表
func (qs *Qstat) SetAttribs(attribs []utils.Attrib) {
	qs.Attribs = attribs
}

//设定扩展信息列表.
func (qs *Qstat) SetExtend(extend string) {
	qs.Extend = extend
}

//设定Id值
func (qs *Qstat) SetId(id string) {
	qs.Id = id
}

//创建一个新的连接
func (qs *Qstat) ConnectPBS() error {
	var err error
	qs.Handle, err = utils.Pbs_connect(qs.Server)
	if err != nil {
		return errors.NewBadRequest(err, "Cann't connect PBSpro Server")
	}

	return nil
}

//断开连接
func (qs *Qstat) DisconnectPBS() error {
	err := utils.Pbs_disconnect(qs.Handle)
	if err != nil {
		return errors.NewBadRequest(err, "Can't disconnect PBSpro Server")
	}
	return nil
}

func Pbs_attrib2attribl(attribs []utils.Attrib) *C.struct_attrl {
	// Empty array returns null pointer
	if len(attribs) == 0 {
		return nil
	}

	first := &C.struct_attrl{
		value:    C.CString(attribs[0].Value),
		resource: C.CString(attribs[0].Resource),
		name:     C.CString(attribs[0].Name),
		op:       uint32(attribs[0].Op),
	}
	tail := first

	for _, attr := range attribs[1:len(attribs)] {
		tail.next = &C.struct_attrl{
			value:    C.CString(attr.Value),
			resource: C.CString(attr.Resource),
			name:     C.CString(attr.Name),
			op:       uint32(attribs[0].Op),
		}
	}

	return first
}

func Pbs_freeattribl(attrl *C.struct_attrl) {
	for p := attrl; p != nil; p = p.next {
		C.free(unsafe.Pointer(p.name))
		C.free(unsafe.Pointer(p.value))
		C.free(unsafe.Pointer(p.resource))
	}
}

//查询指定作业的信息
func (qs *Qstat) Pbs_statjob() ([]utils.BatchStatus, error) {
	i := C.CString(qs.Id)
	defer C.free(unsafe.Pointer(i))

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	batch_status := C.pbs_statjob(C.int(qs.Handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

//查询指定节点状态
func (qs *Qstat) Pbs_statnode() ([]utils.BatchStatus, error) {
	i := C.CString(qs.Id)
	defer C.free(unsafe.Pointer(i))

	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statnode(C.int(qs.Handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

//查询指定队列信息
func (qs *Qstat) Pbs_statque() ([]utils.BatchStatus, error) {
	i := C.CString(qs.Id)
	defer C.free(unsafe.Pointer(i))

	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statque(C.int(qs.Handle), i, a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

//查询服务信息
func (qs *Qstat) Pbs_statserver() ([]utils.BatchStatus, error) {
	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_statserver(C.int(qs.Handle), a, e)

	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)

	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

//返回JOBID列表
func (qs *Qstat) Pbs_selstat() ([]utils.BatchStatus, error) {
	a := Pbs_attrib2attribl(qs.Attribs)
	defer Pbs_freeattribl(a)

	e := C.CString(qs.Extend)
	defer C.free(unsafe.Pointer(e))

	batch_status := C.pbs_selstat(C.int(qs.Handle), (*C.struct_attropl)(unsafe.Pointer(a)), a, e)

	// FIXME: nil also indicates no jobs matched selection criteria...
	if batch_status == nil {
		return nil, errors.New(utils.Pbs_strerror(int(C.pbs_errno)))
	}
	defer C.pbs_statfree(batch_status)
	batch := get_pbs_batch_status(batch_status)

	return batch, nil
}

//获取信息
func get_pbs_batch_status(batch_status *_Ctype_struct_batch_status) (batch []utils.BatchStatus) {

	for batch_status != nil {
		temp := []utils.Attrib{}
		for attr := batch_status.attribs; attr != nil; attr = attr.next {
			temp = append(temp, utils.Attrib{
				Name:     C.GoString(attr.name),
				Resource: C.GoString(attr.resource),
				Value:    C.GoString(attr.value),
			})
		}

		batch = append(batch, utils.BatchStatus{
			Name:       C.GoString(batch_status.name),
			Text:       C.GoString(batch_status.text),
			Attributes: temp,
		})

		batch_status = batch_status.next
	}
	return batch
}
