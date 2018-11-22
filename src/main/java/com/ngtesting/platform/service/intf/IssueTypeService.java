package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuType;

import java.util.List;

public interface IssueTypeService extends BaseService {

    List<IsuType> list(Integer orgId);

    List<IsuType> list(Integer orgId, Integer prjId);

    List<IsuType> listBySolutionId(Integer solutionId, Integer orgId);
    List<IsuType> listNotInSolution(Integer solutionId, Integer orgId);

    IsuType get(Integer id, Integer orgId);

    IsuType save(IsuType vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    Boolean setDefault(Integer id, Integer orgId);

    Boolean changeOrder(Integer id, String act, Integer orgId);
}
