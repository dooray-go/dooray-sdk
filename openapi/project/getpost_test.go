package project

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/dooray-go/dooray-sdk/openapi/model/project"
)

const getPostSuccessBody = `{
    "header": {
        "isSuccessful": true,
        "resultCode": 0,
        "resultMessage": ""
    },
    "result": {
       "id": "post-1",
       "subject": "업무 상세",
       "project": {
           "id": "proj-1",
           "code": "SDK"
       },
       "taskNumber": "SDK/1",
       "closed": false,
       "createdAt": "2026-09-15T10:00:00+09:00",
       "dueDate": "2026-09-16T18:00:00+09:00",
       "dueDateFlag": true,
       "updatedAt": "2026-09-15T11:00:00+09:00",
       "number": 1,
       "priority": "normal",
       "parent": {
           "id": "parent-1",
           "number": 2,
           "subject": "상위 업무"
       },
       "workflowClass": "registered",
       "workflow": {
           "id": "1",
           "name": "등록"
       },
       "milestone": {
           "id": "1",
           "name": "단계"
       },
       "tags": [{
           "id": "tag-1"
       }],
       "body": {
           "mimeType": "text/x-markdown",
           "content": "new body"
       },
       "users": {
           "from": {
               "type": "member",
               "member": {
                   "organizationmemberid": "from-1"
               }
           },
           "to": [{
               "type": "member",
               "member": {
                   "organizationMemberId": "to-1"
               },
               "workflow": {
                   "id": "1",
                   "name": "등록"
               }
           }],
           "cc": []
       },
      "files": [{
        "id": "file-1",
        "name": "spec.md",
        "size": 128
      }]
    }
}`

func TestGetPost(t *testing.T) {
	const (
		projectID = "proj-1"
		postID    = "post-1"
		apiKey    = "test-api-key"
	)

	var gotMethod, gotAuth, gotProjectID, gotPostID string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /project/v1/projects/{projectId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotAuth = r.Header.Get("Authorization")
		gotProjectID = r.PathValue("projectId")
		gotPostID = r.PathValue("postId")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(getPostSuccessBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	response, err := NewProject(server.URL).GetPost(apiKey, projectID, postID)
	if err != nil {
		t.Fatalf("GetPost returned an error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("method: want %s, got %s", http.MethodGet, gotMethod)
	}
	if gotAuth != "dooray-api "+apiKey {
		t.Errorf("Authorization: want %q, got %q", "dooray-api "+apiKey, gotAuth)
	}
	if gotProjectID != projectID {
		t.Errorf("projectId: want %q, got %q", projectID, gotProjectID)
	}
	if gotPostID != postID {
		t.Errorf("postId: want %q, got %q", postID, gotPostID)
	}
	if response.RawJSON != getPostSuccessBody {
		t.Errorf("RawJSON did not match")
	}
	if response.Result.ID != postID {
		t.Errorf("id: want %q, got %q", postID, response.Result.ID)
	}
	if response.Result.Subject != "업무 상세" {
		t.Errorf("subject: want %q, got %q", "업무 상세", response.Result.Subject)
	}
	if response.Result.Parent.Number != 2 {
		t.Errorf("parent.number: want 2, got %d", response.Result.Parent.Number)
	}
	if response.Result.Body.MimeType != "text/x-markdown" || response.Result.Body.Content != "new body" {
		t.Errorf("body: %#v", response.Result.Body)
	}
	if len(response.Result.Files) != 1 || response.Result.Files[0].ID != "file-1" || response.Result.Files[0].Size != 128 {
		t.Errorf("files: %#v", response.Result.Files)
	}
	if !response.Header.IsSuccessful {
		t.Error("header.isSuccessful: want true, got false")
	}
}

func TestGetPost_SubtaskLiveShape(t *testing.T) {
	const (
		projectID = "1944071709723312151"
		postID    = "4422122034836030569"
	)
	const body = `{
  "header": {
    "resultCode": 0,
    "resultMessage": "",
    "isSuccessful": true
  },
  "result": {
    "id": "4422122034836030569",
    "subject": "가나다라마바사아",
    "project": {
      "id": "1944071709723312151",
      "code": "만티스쿠바-개발",
      "projectCategoryId": null
    },
    "taskNumber": "만티스쿠바-개발/4",
    "closed": false,
    "createdAt": "2026-09-15T08:47:06Z",
    "dueDateFlag": true,
    "updatedAt": "2026-09-15T08:47:06Z",
    "number": 4,
    "priority": "none",
    "parent": {
      "id": "2416542752411144246",
      "number": 2,
      "subject": "Kotlin 앱 배포"
    },
    "tags": [],
    "body": {
      "mimeType": "text/html",
      "content": "<div>가나다라마바사아</div>"
    },
    "users": {
      "from": {
        "type": "member",
        "member": {
          "organizationMemberId": "1900555447797833080",
          "name": "작성자"
        }
      },
      "to": [],
      "cc": []
    },
    "fileIdList": [],
    "workflowClass": "registered",
    "milestone": null,
    "workflow": {
      "id": "1944071710042819264",
      "name": "등록"
    }
  }
}`

	mux := http.NewServeMux()
	mux.HandleFunc("GET /project/v1/projects/{projectId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	response, err := NewProject(server.URL).GetPost("key", projectID, postID)
	if err != nil {
		t.Fatalf("GetPost: %v", err)
	}
	if response.Result.ID != postID {
		t.Errorf("id: want %q, got %q", postID, response.Result.ID)
	}
	if response.Result.Body.MimeType != "text/html" || response.Result.Body.Content == "" {
		t.Errorf("body not parsed: %#v", response.Result.Body)
	}
	if response.Result.Parent.Number != 2 {
		t.Errorf("parent.number: want 2 (JSON number), got %d", response.Result.Parent.Number)
	}
	if response.Result.Parent.ID != "2416542752411144246" {
		t.Errorf("parent.id: got %q", response.Result.Parent.ID)
	}
	if response.Result.Users.From.Member.OrganizationMemberID != "1900555447797833080" {
		t.Errorf("from.member.organizationMemberId: got %q", response.Result.Users.From.Member.OrganizationMemberID)
	}
	if response.Result.FileIDList == nil {
		t.Error("fileIdList: want empty slice, got nil")
	}
	if response.Result.Milestone.ID != "" {
		t.Errorf("null milestone should unmarshal to zero value, got %#v", response.Result.Milestone)
	}
}

func TestOrganizationMember_AcceptsCamelAndLegacyJSON(t *testing.T) {
	var camel, legacy model.OrganizationMember
	if err := json.Unmarshal([]byte(`{"organizationMemberId":"camel","name":"n"}`), &camel); err != nil {
		t.Fatal(err)
	}
	if camel.OrganizationMemberID != "camel" || camel.Name != "n" {
		t.Errorf("camel: %#v", camel)
	}
	if err := json.Unmarshal([]byte(`{"organizationmemberid":"legacy"}`), &legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.OrganizationMemberID != "legacy" {
		t.Errorf("legacy: %#v", legacy)
	}
}

func TestGetPost_NotFound(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /project/v1/projects/{projectId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":false,"resultCode":404,"resultMessage":"not found"}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	_, err := NewProject(server.URL).GetPost("key", "missing-project", "missing-post")
	if err == nil {
		t.Fatal("expected error for 404, got nil")
	}
}
