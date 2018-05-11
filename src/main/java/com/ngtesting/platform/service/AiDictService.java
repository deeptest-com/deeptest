package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.AiDict;

import java.util.List;
import java.util.Map;

public interface AiDictService extends BaseService {
	void dictPers();
	AiDict find(String category, String phrase);
	Map<String, List<String>> get(String json);

	List<String> genRegDict(String dict);
}
