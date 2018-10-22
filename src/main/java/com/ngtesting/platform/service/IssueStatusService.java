package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuStatus;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

public interface IssueStatusService extends BaseService {
    List<IsuStatus> list(Integer orgId);

    List<IsuStatus> list(Integer orgId, Integer prjId);

    IsuStatus get(Integer id, Integer orgId);

    IsuStatus save(IsuStatus vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    @Transactional
    Boolean setDefault(Integer id, Integer orgId);

    @Transactional
    Boolean changeOrder(Integer id, String act, Integer orgId);
}
