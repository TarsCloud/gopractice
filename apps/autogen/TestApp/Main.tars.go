// Package TestApp comment
// This file was generated by tars2go 1.1.8
// Generated from proto.tars
package TestApp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	"github.com/TarsCloud/TarsGo/tars/util/trace"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = codec.FromInt8
	_ = unsafe.Pointer(nil)
	_ = bytes.ErrTooLarge
)

// Main struct
type Main struct {
	servant m.Servant
}

// Add is the proxy function for the method defined in the tars file, with the context
func (obj *Main) Add(req *AddReq, rsp *AddRsp, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}
	tarsResp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = obj.servant.TarsInvoke(tarsCtx, 0, "Add", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(tarsResp.SBuffer))
	err = (*rsp).ReadBlock(readBuf, 2, true)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// AddWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Main) AddWithContext(tarsCtx context.Context, req *AddReq, rsp *AddRsp, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	traceData, ok := current.GetTraceData(tarsCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		tars.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, tars.GetClientConfig().ModuleName, obj.servant.Name(), "Add", 0, traceParam, "")
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 0, "Add", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(tarsResp.SBuffer))
	err = (*rsp).ReadBlock(readBuf, 2, true)
	if err != nil {
		return err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["rsp"] = *rsp
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		tars.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, tars.GetClientConfig().ModuleName, obj.servant.Name(), "Add", int(tarsResp.IRet), traceParam, "")
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// AddOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Main) AddOneWayWithContext(tarsCtx context.Context, req *AddReq, rsp *AddRsp, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 1, "Add", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// Sub is the proxy function for the method defined in the tars file, with the context
func (obj *Main) Sub(req *SubReq, rsp *SubRsp, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}
	tarsResp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = obj.servant.TarsInvoke(tarsCtx, 0, "Sub", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(tarsResp.SBuffer))
	err = (*rsp).ReadBlock(readBuf, 2, true)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// SubWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Main) SubWithContext(tarsCtx context.Context, req *SubReq, rsp *SubRsp, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	traceData, ok := current.GetTraceData(tarsCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		tars.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, tars.GetClientConfig().ModuleName, obj.servant.Name(), "Sub", 0, traceParam, "")
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 0, "Sub", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(tarsResp.SBuffer))
	err = (*rsp).ReadBlock(readBuf, 2, true)
	if err != nil {
		return err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["rsp"] = *rsp
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		tars.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, tars.GetClientConfig().ModuleName, obj.servant.Name(), "Sub", int(tarsResp.IRet), traceParam, "")
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// SubOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Main) SubOneWayWithContext(tarsCtx context.Context, req *SubReq, rsp *SubRsp, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = (*rsp).WriteBlock(buf, 2)
	if err != nil {
		return err
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 1, "Sub", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// SetServant sets servant for the service.
func (obj *Main) SetServant(servant m.Servant) {
	obj.servant = servant
}

// TarsSetTimeout sets the timeout for the servant which is in ms.
func (obj *Main) TarsSetTimeout(timeout int) {
	obj.servant.TarsSetTimeout(timeout)
}

// TarsSetProtocol sets the protocol for the servant.
func (obj *Main) TarsSetProtocol(p m.Protocol) {
	obj.servant.TarsSetProtocol(p)
}

// AddServant adds servant  for the service.
func (obj *Main) AddServant(imp MainServant, servantObj string) {
	tars.AddServant(obj, imp, servantObj)
}

// AddServantWithContext adds servant  for the service with context.
func (obj *Main) AddServantWithContext(imp MainServantWithContext, servantObj string) {
	tars.AddServantWithContext(obj, imp, servantObj)
}

type MainServant interface {
	Add(req *AddReq, rsp *AddRsp) (err error)
	Sub(req *SubReq, rsp *SubRsp) (err error)
}
type MainServantWithContext interface {
	Add(tarsCtx context.Context, req *AddReq, rsp *AddRsp) (err error)
	Sub(tarsCtx context.Context, req *SubReq, rsp *SubRsp) (err error)
}

// Dispatch is used to call the server side implement for the method defined in the tars file. withContext shows using context or not.
func (obj *Main) Dispatch(tarsCtx context.Context, val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	buf := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "Add":
		var req AddReq
		var rsp AddRsp

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(readBuf, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			reqTup := tup.NewUniAttribute()
			reqTup.Decode(readBuf)

			var tupBuffer []byte

			reqTup.GetBuffer("req", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = req.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["req"])
				req.ResetDefault()
				if err = json.Unmarshal(jsonStr, &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(tarsCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			tars.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, tars.GetClientConfig().ModuleName, tarsReq.SServantName, "Add", 0, traceParam, "")
		}

		if !withContext {
			imp := val.(MainServant)
			err = imp.Add(&req, &rsp)
		} else {
			imp := val.(MainServantWithContext)
			err = imp.Add(tarsCtx, &req, &rsp)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			buf.Reset()

			err = rsp.WriteBlock(buf, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			rspTup := tup.NewUniAttribute()

			buf.Reset()
			err = rsp.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

			rspTup.PutBuffer("rsp", buf.ToBytes())

			buf.Reset()
			err = rspTup.Encode(buf)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["rsp"] = rsp

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}

		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSS, uint(buf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["rsp"] = rsp
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			tars.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, tars.GetClientConfig().ModuleName, tarsReq.SServantName, "Add", 0, traceParam, "")
		}

	case "Sub":
		var req SubReq
		var rsp SubRsp

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(readBuf, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			reqTup := tup.NewUniAttribute()
			reqTup.Decode(readBuf)

			var tupBuffer []byte

			reqTup.GetBuffer("req", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = req.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["req"])
				req.ResetDefault()
				if err = json.Unmarshal(jsonStr, &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(tarsCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			tars.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, tars.GetClientConfig().ModuleName, tarsReq.SServantName, "Sub", 0, traceParam, "")
		}

		if !withContext {
			imp := val.(MainServant)
			err = imp.Sub(&req, &rsp)
		} else {
			imp := val.(MainServantWithContext)
			err = imp.Sub(tarsCtx, &req, &rsp)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			buf.Reset()

			err = rsp.WriteBlock(buf, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			rspTup := tup.NewUniAttribute()

			buf.Reset()
			err = rsp.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

			rspTup.PutBuffer("rsp", buf.ToBytes())

			buf.Reset()
			err = rspTup.Encode(buf)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["rsp"] = rsp

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}

		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSS, uint(buf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["rsp"] = rsp
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			tars.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, tars.GetClientConfig().ModuleName, tarsReq.SServantName, "Sub", 0, traceParam, "")
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var statusMap map[string]string
	if status, ok := current.GetResponseStatus(tarsCtx); ok && status != nil {
		statusMap = status
	}
	var contextMap map[string]string
	if ctx, ok := current.GetResponseContext(tarsCtx); ok && ctx != nil {
		contextMap = ctx
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(buf.ToBytes()),
		Status:       statusMap,
		SResultDesc:  "",
		Context:      contextMap,
	}

	_ = readBuf
	_ = buf
	_ = length
	_ = have
	_ = ty
	return nil
}