package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuResolution;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

public interface IssueResolutionService extends BaseService {
    List<IsuResolution> list(Integer orgId);

    List<IsuResolution> list(Integer orgId, Integer prjId);

    IsuResolution get(Integer id, Integer orgId);

    IsuResolution save(IsuResolution vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    @Transactional
    Boolean setDefault(Integer id, Integer orgId);

    @Transactional
    Boolean changeOrder(Integer id, String act, Integer orgId);
}
