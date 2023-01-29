namespace go api

service AppMine {
    ReviewProjectListResponse ReviewProjectList(1: ReviewProjectListParams req)
    ReviewProjectDetailResponse ReviewProjectDetails(1: IdsInt64Params req)
    SaveResponse ReviewProjectSave(1: ReviewProjectSaveParam req)
    SaveResponse ReviewProjectStatus(1: StatusParam req)
    SaveResponse ReviewProjectDelete(1: IdParam req)
}
struct IdsParams {
	1: list<i32> ids
}
struct IdsInt64Params {
	1: list<i64> ids
}
struct StrIdParam {
	1: string id
}
struct IdParam {
	1: i64 id
}
struct SaveResponse {
	1: i64 rowsAffected
}
struct StatusParam {
    1: i64      id
    2: i8       status
    3: string   remark
}
struct ReviewProjectListParams {
    1: i32      pageNum
    2: i32      pageSize
    3: i64      id
    4: string   name
    5: string   modeCode
    6: i8       status
    7: i8       showStatus
}
struct ReviewProjectData  {
	1:  i64     id
	2:  string  name
	3:  string  modeCode
	4:  string  description
	5:  i8      status
	6:  i8      deleted
	7:  string  createdUser
    8:  string  updatedUser
    9:  string  createdDate
    10: string  updatedDate
}
struct ReviewProjectSaveData  {
	1:  i64     id
	2:  string  name
	3:  string  modeCode
	4:  string  description
}
struct ReviewProjectListResponse {
   1: i64           total
   2: i32           pageNum
   3: i32           pageSize
   4: list<ReviewProjectData>    data
}
struct ReviewProjectDetailResponse {
	1: list<ReviewProjectData> data
}
struct ReviewProjectSaveParam {
    1: string           action
    2: ReviewProjectSaveData   data
}

