# 接口层

​接口适配器层​：负责将外部输入（如 HTTP 请求）转换为内部可用的格式（如 DTO），再将内部结果转换为外部输出（如 JSON 响应。
​DTO 应靠近接口层，避免污染领域层或应用层。