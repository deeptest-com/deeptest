package cn.linkr.testspace.service;

import org.springframework.stereotype.Service;

import cn.linkr.testspace.vo.JsonBean;
import cn.linkr.testspace.vo.JsonResult;

@Service
public interface HttpClientService extends BaseService {

    /**
     * 发送 post请求访问本地应用并根据传递参数不同返回不同结果
     *
     * @param url  请求路径
     * @param json 参数json字符串
     */
    public String post(String url, String json);

}
