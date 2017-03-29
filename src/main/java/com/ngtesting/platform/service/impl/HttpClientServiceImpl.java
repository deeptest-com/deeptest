package com.ngtesting.platform.service.impl;

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
import com.ngtesting.platform.service.HttpClientService;
import com.ngtesting.platform.vo.JsonBean;
import com.ngtesting.platform.vo.JsonResult;

import com.alibaba.fastjson.JSONObject;

@Service
public class HttpClientServiceImpl extends BaseServiceImpl implements HttpClientService {
    Log logger = LogFactory.getLog(HttpClientServiceImpl.class);
    
    public static final String PROTOCOL_HTTP = "http://";
    public static final String PROTOCOL_HTTPS = "https://";
    public static final String ContentType = "application/json";
    public static final String Encoding = "UTF-8";

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
            StringEntity entity = new StringEntity(json, Encoding);
            entity.setContentEncoding(Encoding);
            entity.setContentType(ContentType);
            httppost.setEntity(entity);
            response = httpclient.execute(httppost);
            HttpEntity result = response.getEntity();
            resultJson = EntityUtils.toString(result, Encoding);
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
