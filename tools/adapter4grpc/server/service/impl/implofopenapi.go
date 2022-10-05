package impl

import (
	"context"
	"fstools/adapter4grpc/server/dto"
	"fstools/adapter4grpc/server/service"

	"github.com/wup364/filestorage/opensdk"
)

// OpenApiServerImpl implementations.
type OpenApiServerImpl struct {
	sdk opensdk.IOpenApi
}

func (o *OpenApiServerImpl) IsDir(ctx context.Context, dto *dto.QryOfString) (res *dto.ResultOfBool, err error) {
	res.Result, err = o.sdk.IsDir(dto.Query)
	return
}
func (o *OpenApiServerImpl) IsFile(ctx context.Context, dto *dto.QryOfString) (res *dto.ResultOfBool, err error) {
	res.Result, err = o.sdk.IsFile(dto.Query)
	return
}
func (o *OpenApiServerImpl) IsExist(ctx context.Context, dto *dto.QryOfString) (res *dto.ResultOfBool, err error) {
	res.Result, err = o.sdk.IsExist(dto.Query)
	return
}
func (o *OpenApiServerImpl) GetFileSize(ctx context.Context, dto *dto.QryOfString) (res *dto.ResultOfInt64, err error) {
	res.Result, err = o.sdk.GetFileSize(dto.Query)
	return
}
func (o *OpenApiServerImpl) GetNode(ctx context.Context, dto *dto.QryOfString) (res *service.TNode, err error) {
	var node *opensdk.TNode
	if node, err = o.sdk.GetNode(dto.Query); nil == err {
		res = &service.TNode{
			Id:    node.Id,
			Pid:   node.Pid,
			Addr:  node.Addr,
			Flag:  int32(node.Flag),
			Name:  node.Name,
			Size:  node.Size,
			Ctime: node.Ctime,
			Mtime: node.Mtime,
			Props: node.Props,
		}
	}
	return
}
func (o *OpenApiServerImpl) GetNodes(ctx context.Context, dto *dto.QryOfStrings) (res *service.DirNodeListDto, err error) {
	var nodes []opensdk.TNode
	if nodes, err = o.sdk.GetNodes(dto.Query); nil == err && len(nodes) > 0 {
		res.Total = int64(len(nodes))
		res.Datas = make([]*service.TNode, res.Total)
		for i := int64(0); i < res.Total; i++ {
			res.Datas[i] = &service.TNode{
				Id:    nodes[i].Id,
				Pid:   nodes[i].Pid,
				Addr:  nodes[i].Addr,
				Flag:  int32(nodes[i].Flag),
				Name:  nodes[i].Name,
				Size:  nodes[i].Size,
				Ctime: nodes[i].Ctime,
				Mtime: nodes[i].Mtime,
				Props: nodes[i].Props,
			}
		}
	}
	return
}
func (o *OpenApiServerImpl) GetDirNameList(ctx context.Context, dto *dto.QueryLimitOfString) (res *service.DirNameListDto, err error) {
	var nodes *opensdk.DirNameListDto
	if nodes, err = o.sdk.GetDirNameList(dto.Query, int(dto.Limit), int(dto.Offset)); nil == err && nodes.Total > 0 {
		res.Total = int64(nodes.Total)
		res.Datas = nodes.Datas
	}
	return
}
func (o *OpenApiServerImpl) GetDirNodeList(ctx context.Context, dto *dto.QueryLimitOfString) (res *service.DirNodeListDto, err error) {
	var nodes *opensdk.DirNodeListDto
	if nodes, err = o.sdk.GetDirNodeList(dto.Query, int(dto.Limit), int(dto.Offset)); nil == err && nodes.Total > 0 {
		res.Total = int64(nodes.Total)
		res.Datas = make([]*service.TNode, res.Total)
		for i := int64(0); i < res.Total; i++ {
			res.Datas[i] = &service.TNode{
				Id:    nodes.Datas[i].Id,
				Pid:   nodes.Datas[i].Pid,
				Addr:  nodes.Datas[i].Addr,
				Flag:  int32(nodes.Datas[i].Flag),
				Name:  nodes.Datas[i].Name,
				Size:  nodes.Datas[i].Size,
				Ctime: nodes.Datas[i].Ctime,
				Mtime: nodes.Datas[i].Mtime,
				Props: nodes.Datas[i].Props,
			}
		}
	}
	return
}
func (o *OpenApiServerImpl) DoMkDir(ctx context.Context, dto *dto.QryOfString) (res *dto.ResultOfString, err error) {
	res.Result, err = o.sdk.DoMkDir(dto.Query)
	return
}
func (o *OpenApiServerImpl) DoDelete(ctx context.Context, dto *dto.QryOfString) (res *dto.ResultOfBool, err error) {
	if err = o.sdk.DoDelete(dto.Query); nil == err {
		res.Result = true
	}
	return
}
func (o *OpenApiServerImpl) DoRename(ctx context.Context, dto *service.RenameCmd) (res *dto.ResultOfBool, err error) {
	if err = o.sdk.DoRename(dto.Src, dto.Dst); nil == err {
		res.Result = true
	}
	return
}
func (o *OpenApiServerImpl) DoCopy(ctx context.Context, dto *service.MoveCmd) (res *dto.ResultOfString, err error) {
	res.Result, err = o.sdk.DoCopy(dto.Src, dto.Dst, dto.Override)
	return
}
func (o *OpenApiServerImpl) DoMove(ctx context.Context, dto *service.CopyCmd) (res *dto.ResultOfBool, err error) {
	if err = o.sdk.DoMove(dto.Src, dto.Dst, dto.Override); nil == err {
		res.Result = true
	}
	return
}
func (o *OpenApiServerImpl) DoQueryToken(ctx context.Context, dto *dto.QryOfString) (res *service.StreamToken, err error) {
	var tmp *opensdk.StreamToken
	if tmp, err = o.sdk.DoQueryToken(dto.Query); nil == err {
		res = &service.StreamToken{
			Token:    tmp.Token,
			NodeNo:   tmp.NodeNo,
			FileID:   tmp.FileID,
			FilePath: tmp.FilePath,
			FileSize: tmp.FileSize,
			CTime:    tmp.CTime,
			MTime:    tmp.MTime,
			EndPoint: tmp.EndPoint,
			Type:     int32(tmp.Type),
		}
	}
	return
}
func (o *OpenApiServerImpl) DoAskReadToken(ctx context.Context, dto *dto.QryOfString) (res *service.StreamToken, err error) {
	var tmp *opensdk.StreamToken
	if tmp, err = o.sdk.DoAskReadToken(dto.Query); nil == err {
		res = &service.StreamToken{
			Token:    tmp.Token,
			NodeNo:   tmp.NodeNo,
			FileID:   tmp.FileID,
			FilePath: tmp.FilePath,
			FileSize: tmp.FileSize,
			CTime:    tmp.CTime,
			MTime:    tmp.MTime,
			EndPoint: tmp.EndPoint,
			Type:     int32(tmp.Type),
		}
	}
	return
}
func (o *OpenApiServerImpl) DoAskWriteToken(ctx context.Context, dto *dto.QryOfString) (res *service.StreamToken, err error) {
	var tmp *opensdk.StreamToken
	if tmp, err = o.sdk.DoAskWriteToken(dto.Query); nil == err {
		res = &service.StreamToken{
			Token:    tmp.Token,
			NodeNo:   tmp.NodeNo,
			FileID:   tmp.FileID,
			FilePath: tmp.FilePath,
			FileSize: tmp.FileSize,
			CTime:    tmp.CTime,
			MTime:    tmp.MTime,
			EndPoint: tmp.EndPoint,
			Type:     int32(tmp.Type),
		}
	}
	return
}
func (o *OpenApiServerImpl) DoRefreshToken(ctx context.Context, dto *dto.QryOfString) (res *service.StreamToken, err error) {
	var tmp *opensdk.StreamToken
	if tmp, err = o.sdk.DoRefreshToken(dto.Query); nil == err {
		res = &service.StreamToken{
			Token:    tmp.Token,
			NodeNo:   tmp.NodeNo,
			FileID:   tmp.FileID,
			FilePath: tmp.FilePath,
			FileSize: tmp.FileSize,
			CTime:    tmp.CTime,
			MTime:    tmp.MTime,
			EndPoint: tmp.EndPoint,
			Type:     int32(tmp.Type),
		}
	}
	return
}
func (o *OpenApiServerImpl) DoSubmitWriteToken(ctx context.Context, dto *service.SubmitTokenCmd) (res *service.TNode, err error) {
	var tmp *opensdk.TNode
	if tmp, err = o.sdk.DoSubmitWriteToken(dto.Token, dto.Override); nil == err {
		res = &service.TNode{
			Id:    tmp.Id,
			Pid:   tmp.Pid,
			Addr:  tmp.Addr,
			Flag:  int32(tmp.Flag),
			Name:  tmp.Name,
			Size:  tmp.Size,
			Ctime: tmp.Ctime,
			Mtime: tmp.Mtime,
			Props: tmp.Props,
		}
	}
	return
}
func (o *OpenApiServerImpl) GetReadStreamURL(ctx context.Context, dto *service.QryStreamURLCmd) (res *dto.ResultOfString, err error) {
	res.Result, err = o.sdk.GetReadStreamURL(dto.NodeNo, dto.Token, dto.Endpoint)
	return
}
func (o *OpenApiServerImpl) GetWriteStreamURL(ctx context.Context, dto *service.QryStreamURLCmd) (res *dto.ResultOfString, err error) {
	res.Result, err = o.sdk.GetWriteStreamURL(dto.NodeNo, dto.Token, dto.Endpoint)
	return
}
