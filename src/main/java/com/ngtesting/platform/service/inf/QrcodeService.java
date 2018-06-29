package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;

public interface QrcodeService extends BaseService {

	JSONObject decode(String filePath);

}
