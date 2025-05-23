package transform

import (
	"context"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog/v2"
	transformersv1alpha1 "k8s.io/sample-apiserver/pkg/apis/transformers/v1alpha1"
)

// 子资源只需要实现这几个接口，主要是 Connecter 接口
var (
	_ rest.Connecter = &REST{}
	_ rest.Storage   = &REST{}
	_ rest.Scoper    = &REST{}
)

// REST 实现了 transform 子资源的处理
type REST struct{}

// NamespaceScoped 返回此资源是否在命名空间范围内
func (r *REST) NamespaceScoped() bool {
	return true
}

// Destroy 实现 rest.Storage 接口
func (r *REST) Destroy() {
	// 无需清理资源
}

// New 实现 rest.Storage 接口
func (r *REST) New() runtime.Object {
	// 这个子资源不需要存储，所以随便返回一个
	return &transformersv1alpha1.Autobot{}
}

// Connect 实现子资源的实际逻辑处理
func (r *REST) Connect(ctx context.Context, name string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
	klog.Infof("Transform 子资源 Connect 被调用：资源名=%s", name)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		// 可以通过 name 获取 autobot 对象，执行实际的转换逻辑
		// autobot, err := getAutobotClient().Get(ctx, name, metav1.GetOptions{})

		if req.Method == http.MethodPost {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Autobot '%s' 转换完成！(POST方法)", name)))
		} else if req.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Autobot '%s' 转换状态查询 (GET方法)", name)))
		} else {
			http.Error(w, fmt.Sprintf("不支持的HTTP方法 %s", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}), nil
}

// ConnectMethods 返回此子资源支持的HTTP方法
func (r *REST) ConnectMethods() []string {
	return []string{"GET", "POST"}
}

// NewConnectOptions 创建用于连接操作的选项对象
func (r *REST) NewConnectOptions() (runtime.Object, bool, string) {
	// 第一个返回值：选项对象（如果不需要特殊选项，返回nil）
	// 第二个返回值：是否需要重定向
	// 第三个返回值：重定向路径（如果需要）
	return nil, false, ""
}

// NewREST 创建一个新的 REST 实例
func NewREST() *REST {
	return &REST{}
}
