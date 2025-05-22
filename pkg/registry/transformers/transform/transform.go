package transformers

// TransformerREST 实现了一个 REST 存储，用于处理 optimusprime 资源的变形子资源
// type TransformerREST struct {
// 	store rest.StandardStorage
// }

// // 确保 TransformerREST 实现了 rest.Connecter 接口
// var _ rest.Connecter = &TransformerREST{}

// // 确保 TransformerREST 实现了 rest.Storage 接口
// var _ rest.Storage = &TransformerREST{}

// // NewTransformerREST 创建 TransformerREST 实例
// func NewTransformerREST(store rest.StandardStorage) *TransformerREST {
// 	return &TransformerREST{
// 		store: store,
// 	}
// }

// // New 返回一个空的 TransformRequest 对象
// func (r *TransformerREST) New() runtime.Object {
// 	return &TransformRequest{}
// }

// // Destroy 实现 rest.Storage 接口
// func (r *TransformerREST) Destroy() {
// 	// 什么都不做，因为我们没有需要清理的资源
// }

// // Connect 处理变形请求
// func (r *TransformerREST) Connect(ctx context.Context, name string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
// 	// 获取变形选项
// 	transformOpts, ok := opts.(*TransformRequest)
// 	if !ok {
// 		return nil, fmt.Errorf("invalid options object: %#v", opts)
// 	}

// 	// 获取 optimusprime 资源对象
// 	obj, err := r.store.Get(ctx, name, &metav1.GetOptions{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 将对象转换为 Fischer (在实际代码中，应该是转换为 OptimusPrime 类型)
// 	_, ok = obj.(*transformers.OptimusPrime)
// 	if !ok {
// 		return nil, fmt.Errorf("不是有效的 OptimusPrime 对象")
// 	}

// 	// 创建 HTTP handler 处理请求
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		// 只支持 POST 方法
// 		if req.Method != "POST" {
// 			responder.Error(fmt.Errorf("只支持 POST, 收到了 %s", req.Method))
// 			return
// 		}

// 		// 根据变形模式设置响应
// 		var form string
// 		switch transformOpts.Mode {
// 		case "robot":
// 			form = "Robot Form"
// 		case "vehicle":
// 			form = "Vehicle Form"
// 		case "weapon":
// 			form = "Weapon Form"
// 		default:
// 			form = "Unknown Form"
// 		}

// 		// 这里你可以根据需要更新 optimusprime 资源的状态
// 		// 例如可以通过 r.store.Update 更新资源

// 		// 返回变形响应
// 		response := &TransformResponse{
// 			Status: "变形成功",
// 			Form:   form,
// 		}
// 		response.TypeMeta.APIVersion = transformOpts.APIVersion
// 		response.TypeMeta.Kind = "TransformResponse"

// 		responder.Object(http.StatusOK, response)
// 	}), nil
// }

// // NewConnectOptions 返回一个新的变形请求选项
// func (r *TransformerREST) NewConnectOptions() (runtime.Object, bool, string) {
// 	return &TransformRequest{}, false, "mode"
// }

// // ConnectMethods 返回支持的 HTTP 方法
// func (r *TransformerREST) ConnectMethods() []string {
// 	return []string{"POST"}
// }
