package cn.linkr.events.service;

import org.springframework.stereotype.Service;

import cn.linkr.events.vo.JsonBean;
import cn.linkr.events.vo.JsonResult;

@Service
public interface HttpClientService extends BaseService {
    /**
     * 手机操作接口
     *
     * @param deviceId 设备id
     * @param bean     封装了操作类型和参数的对象
     * @return 操作结果
     */
    public JsonResult operator(Long hostIp, String FwPort, JsonBean bean);

    /**
     * 发送 post请求访问本地应用并根据传递参数不同返回不同结果
     *
     * @param url  请求路径
     * @param json 参数json字符串
     */
    public String post(String url, String json);

}
