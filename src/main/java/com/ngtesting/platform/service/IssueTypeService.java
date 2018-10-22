package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuType;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

public interface IssueTypeService extends BaseService {

    List<IsuType> list(Integer orgId);

    List<IsuType> list(Integer orgId, Integer prjId);

    IsuType get(Integer id, Integer orgId);

    IsuType save(IsuType vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    @Transactional
    Boolean setDefault(Integer id, Integer orgId);

    @Transactional
    Boolean changeOrder(Integer id, String act, Integer orgId);
}
