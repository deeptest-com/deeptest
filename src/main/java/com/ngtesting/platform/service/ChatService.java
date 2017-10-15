package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestThread;



public interface ChatService extends BaseService {

	List<TestThread> listByEvent(Long eventId);

	TestThread save(Long eventId, Long parentId, Long clientId, String content);

	List<TestThread> enter(Long eventId, Long clientId);


}
