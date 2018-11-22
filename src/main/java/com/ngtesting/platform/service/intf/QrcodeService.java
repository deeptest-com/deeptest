package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;

public interface QrcodeService extends BaseService {

	JSONObject decode(String filePath);

}
