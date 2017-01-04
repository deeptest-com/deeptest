package cn.mobiu.events.service;

import com.alibaba.fastjson.JSONObject;;

public interface QrcodeService extends BaseService {

	JSONObject decode(String filePath);

}
