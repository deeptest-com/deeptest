package cn.linkr.events.service;

import com.alibaba.fastjson.JSONObject;;

public interface QrcodeService extends BaseService {

	JSONObject decode(String filePath);

}
