package cn.linkr.events.service.impl;

import java.io.IOException;

import org.apache.commons.lang.StringUtils;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.apache.http.HttpEntity;
import org.apache.http.client.config.RequestConfig;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.springframework.stereotype.Service;

import cn.linkr.events.constants.HttpServerConstants;
import cn.linkr.events.service.HttpClientService;
import cn.linkr.events.vo.JsonBean;
import cn.linkr.events.vo.JsonResult;

import com.alibaba.fastjson.JSONObject;

@Service
public class HttpClientServiceImpl extends BaseServiceImpl implements HttpClientService {
    Log logger = LogFactory.getLog(HttpClientServiceImpl.class);

    @Override
    public JsonResult operator(Long hostIp, String FwPort, JsonBean bean) {

        String url = HttpServerConstants.PROTOCOL_HTTP + hostIp + ":" +
                (HttpServerConstants.ServicePortPrefix + FwPort);
        String json = JSONObject.toJSONString(bean);
        String result = this.post(url, json);
        JsonResult jsonResult = null;
        if (StringUtils.isNotEmpty(result)) {
            jsonResult = JSONObject.parseObject(result, JsonResult.class);
        }
        return jsonResult;
    }

    @Override
    public String post(String url, String json) {
        String resultJson = "";
        // 创建默认的httpClient实例.
        CloseableHttpClient httpclient = HttpClients.createDefault();

        // 创建httppost
        HttpPost httppost = new HttpPost(url);
        RequestConfig requestConfig = RequestConfig.custom().setSocketTimeout(2000).setConnectTimeout(2000).build();//设置请求和传输超时时间
        httppost.setConfig(requestConfig);
        CloseableHttpResponse response = null;
        try {
            StringEntity entity = new StringEntity(json, HttpServerConstants.Encoding);
            entity.setContentEncoding(HttpServerConstants.Encoding);
            entity.setContentType(HttpServerConstants.ContentType);
            httppost.setEntity(entity);
            response = httpclient.execute(httppost);
            HttpEntity result = response.getEntity();
            resultJson = EntityUtils.toString(result, HttpServerConstants.Encoding);
        } catch (Exception e) {
            logger.error(e.getMessage());
        } finally {
            // 关闭连接,释放资源
            if (response != null) {
                try {
                    response.close();
                } catch (IOException e) {
                }
            }
            try {
                httpclient.close();
            } catch (IOException e) {
            }
        }
        return resultJson;
    }
}
