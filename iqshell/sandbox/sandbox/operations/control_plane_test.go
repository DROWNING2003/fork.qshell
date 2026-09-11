//go:build unit

package operations

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	sbClient "github.com/qiniu/qshell/v2/iqshell/sandbox"
)

func newSandboxOperationServer(t *testing.T, statusCode int, body string, requests *[]string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*requests = append(*requests, fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(body))
	}))
}

func configureSandboxOperationClient(t *testing.T, endpoint string) {
	t.Helper()
	t.Setenv(sbClient.EnvQiniuAPIKey, "test-api-key")
	t.Setenv(sbClient.EnvQiniuSandboxAPIURL, endpoint)
	t.Setenv("SANDBOX_RETRY_MAX", "0")
}

func TestListResourcesUsesResourceAPI(t *testing.T) {
	var requests []string
	server := newSandboxOperationServer(t, http.StatusOK, `{"resources":[]}`, &requests)
	defer server.Close()
	configureSandboxOperationClient(t, server.URL)

	ListResources(ResourceListInfo{SandboxID: "sandbox-1", Format: sbClient.FormatJSON})

	if len(requests) != 1 || requests[0] != "GET /sandboxes/sandbox-1/resources" {
		t.Fatalf("requests = %v, want direct resource GET", requests)
	}
}

func TestUpdateResourceTokenUsesPatch(t *testing.T) {
	var requests []string
	server := newSandboxOperationServer(t, http.StatusNoContent, "", &requests)
	defer server.Close()
	configureSandboxOperationClient(t, server.URL)

	UpdateResourceToken(ResourceUpdateInfo{
		SandboxID:  "sandbox-1",
		ResourceID: "resource-1",
		Token:      "token",
	})

	if len(requests) != 1 || requests[0] != "PATCH /sandboxes/sandbox-1/resources/resource-1" {
		t.Fatalf("requests = %v, want direct resource PATCH", requests)
	}
}

func TestLogsUsesLogsAPI(t *testing.T) {
	var requests []string
	server := newSandboxOperationServer(t, http.StatusOK, `{"logs":[],"logEntries":[]}`, &requests)
	defer server.Close()
	configureSandboxOperationClient(t, server.URL)

	Logs(LogsInfo{SandboxID: "sandbox-1", Format: sbClient.FormatJSON})

	if len(requests) != 1 || requests[0] != "GET /sandboxes/sandbox-1/logs" {
		t.Fatalf("requests = %v, want direct logs GET", requests)
	}
}

func TestMetricsUsesMetricsAPI(t *testing.T) {
	var requests []string
	server := newSandboxOperationServer(t, http.StatusOK, `[]`, &requests)
	defer server.Close()
	configureSandboxOperationClient(t, server.URL)

	Metrics(MetricsInfo{SandboxID: "sandbox-1", Format: sbClient.FormatJSON})

	if len(requests) != 1 || requests[0] != "GET /sandboxes/sandbox-1/metrics" {
		t.Fatalf("requests = %v, want direct metrics GET", requests)
	}
}

func TestKillUsesDeleteAPI(t *testing.T) {
	var requests []string
	server := newSandboxOperationServer(t, http.StatusNoContent, "", &requests)
	defer server.Close()
	configureSandboxOperationClient(t, server.URL)

	Kill(KillInfo{SandboxIDs: []string{"sandbox-1"}})

	if len(requests) != 1 || requests[0] != "DELETE /sandboxes/sandbox-1" {
		t.Fatalf("requests = %v, want direct sandbox DELETE", requests)
	}
}

func TestPauseUsesPauseAPI(t *testing.T) {
	var requests []string
	server := newSandboxOperationServer(t, http.StatusNoContent, "", &requests)
	defer server.Close()
	configureSandboxOperationClient(t, server.URL)

	Pause(PauseInfo{SandboxIDs: []string{"sandbox-1"}})

	if len(requests) != 1 || requests[0] != "POST /sandboxes/sandbox-1/pause" {
		t.Fatalf("requests = %v, want direct sandbox pause", requests)
	}
}
