namespace go subject

include "base.thrift"

struct Subject {
    1: optional i64 ID (api.body="id")
    2: optional string Name (api.body="name")
    3: optional string Answer (api.body="answer")
    4: optional i16 SubjectType (api.body="subject_type")
    5: optional i64 CreatorID (api.body="creator_id")
    6: optional i64 CreatedAt (api.body="created_at")
}

struct BaseSubject {
    1: optional i64 ID (api.body="id")
    2: optional string Name (api.body="name")
    3: optional string Answer (api.body="answer")
    4: optional i16 SubjectType (api.body="subject_type")
}

struct UserWithSubject {
    1: optional i64 UserID (api.body="user_id")
    2: optional i64 SubjectID (api.body="subject_id")
    3: optional i16 SubjectType (api.body="subject_type")
    4: optional i16 Phase (api.body="phase")
    5: optional i16 LearnTimes (api.body="learn_times")
    6: optional i64 LastReviewAt (api.body="last_review_at")
}

struct SubjectRecord {
    1: optional i64 ID (api.body="id")
    2: optional string Name (api.body="name")
    3: optional string Answer (api.body="answer")
    4: optional i16 SubjectType (api.body="subject_type")
    5: optional i16 Phase (api.body="phase")
    6: optional i16 LearnTimes (api.body="learn_times")
    7: optional i64 LastReviewAt (api.body="last_review_at")
    8: optional i64 NextReviewAt (api.body="next_review_at")
}

struct SubjectRecordResp {
    1: optional i64 ID (api.body="id")
    2: optional string Name (api.body="name")
    3: optional string Answer (api.body="answer")
    4: optional i16 SubjectType (api.body="subject_type")
    5: optional i16 Phase (api.body="phase")
    6: optional i16 LearnTimes (api.body="learn_times")
    7: optional string LastReviewAt (api.body="last_review_at")
    8: optional string NextReviewAt (api.body="next_review_at")
}

struct Remind {
    1: optional i64 UserID (api.body="user_id")
    2: optional i64 SubjectID (api.body="subject_id")
    3: optional i64 Remind (api.body="remind")
}

struct GetSubjectsReq {
    1: optional i64 UserID
    2: optional i16 SubjectType (api.query="subject-type",api.path="subject-type")
    3: optional i64 Pages (api.form="pages")
    4: optional i64 Size (api.form="size")
}

struct GetSubjectsResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 Total (api.body="total")
    3: optional list<BaseSubject> Subjects (api.body="subjects")
}

struct AddForgetSubjectReq {
    1: optional i64 UserID
    2: optional i64 SubjectID (api.form="subject_id")
    3: optional i16 SubjectType (api.form="subject_type")
}

struct AddForgetSubjectResp {
    1: optional base.Base Base (api.body="base")
}

struct AddSuccessSubjectReq {
    1: optional i64 UserID
    2: optional i64 SubjectID (api.form="subject_id")
    3: optional i16 SubjectType (api.form="subject_type")
}

struct AddSuccessSubjectResp {
    1: optional base.Base Base (api.body="base")
}

struct GetNeededReviewSubjectsReq {
    1: optional i64 UserID
    2: optional i64 pages (api.form="pages")
    3: optional i64 size (api.form="size")
}

struct GetNeededReviewSubjectsResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 Total (api.body="total")
    3: optional list<SubjectRecordResp> Subjects (api.body="subjects")
}

struct GetAnswerSubjectRecordReq {
    1: optional i64 UserID
    2: optional i64 pages (api.form="pages")
    3: optional i64 size (api.form="size")
}

struct GetAnswerSubjectRecordResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 Total (api.body="total")
    3: optional list<SubjectRecordResp> Subjects (api.body="subjects")
}

struct AddNewSubjectReq {
    1: optional i64 UserID
    2: optional i16 SubjectType (api.form="subject_type")
    3: optional string Name (api.form="name")
    4: optional string Answer (api.form="answer")
}

struct AddNewSubjectResp {
    1: optional base.Base Base (api.body="base")
}

service SubjectService {
    GetSubjectsResp GetSubjects(1: GetSubjectsReq req) (api.post="/api/subject/get")
    AddForgetSubjectResp AddForgetSubject(1: AddForgetSubjectReq req) (api.post="/api/subject/add-forget")
    AddSuccessSubjectResp AddSuccessSubject(1: AddSuccessSubjectReq req) (api.post="/api/subject/add-success")
    GetNeededReviewSubjectsResp GetNeededReviewSubjects(1: GetNeededReviewSubjectsReq req) (api.post="/api/subject/get-needed-review")
    GetAnswerSubjectRecordResp GetAnswerSubjectRecord(1: GetAnswerSubjectRecordReq req) (api.post="/api/subject/get-answer-record")
    AddNewSubjectResp AddNewSubject(1: AddNewSubjectReq req) (api.post="/api/subject/add-new-subject")
}
