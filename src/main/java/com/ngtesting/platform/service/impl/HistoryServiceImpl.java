package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestHistory;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.DateUtils;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestHistoryVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class HistoryServiceImpl extends BaseServiceImpl implements HistoryService {

	@Override
	public List<TestHistory> listByOrg(Long orgId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestHistory.class);

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		dc.createAlias("project", "project");
		dc.add(Restrictions.eq("project.orgId", orgId));

		dc.addOrder(Order.desc("createTime"));

		Page page = findPage(dc, 0, 30);

		return page.getItems();
	}

	@Override
	public List<TestHistory> listByProject(Long projectId, String projectType) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestHistory.class);

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		if (projectType.equals(TestProject.ProjectType.project.toString())) {
			dc.add(Restrictions.eq("projectId", projectId));
		} else {
			dc.createAlias("project", "project");
			dc.add(Restrictions.eq("project.parentId", projectId));
		}

		dc.addOrder(Order.desc("createTime"));

		Page page = findPage(dc, 0, 30);

		return page.getItems();
	}

    @Override
    public TestHistoryVo getById(Long id) {
		TestHistory po = (TestHistory) get(TestHistory.class, id);
		TestHistoryVo vo = genVo(po);

        return vo;
    }
    @Override
    public TestHistory create(Long projectId, UserVo optUser, String action,
                              TestHistory.TargetType entityType, Long entityId, String name) {
        TestHistory history = new TestHistory();

        history.setTitle("用户" + StringUtil.highlightDict(optUser.getName())
                + action + entityType.name + StringUtil.highlightDict(name));
        history.setProjectId(projectId);
        history.setEntityId(entityId);
        history.setEntityType(entityType);
        history.setUserId(optUser.getId());
        saveOrUpdate(history);

        return history;
    }

	@Override
	public Map<String, List<TestHistoryVo>> genVosByDate(List<TestHistory> historyPos) {
		Map<String, List<TestHistoryVo>> map = new LinkedHashMap();
		for(TestHistory his: historyPos) {
            Date createDate = his.getCreateTime();
            String date = DateUtils.FormatDate(createDate, "yyyy-MM-dd");
            if (!map.containsKey(date)) {
                map.put(date, new LinkedList());
            }
            map.get(date).add(genVo(his));
		}
		return map;
	}

	@Override
	public List<TestHistoryVo> genVos(List<TestHistory> pos) {
        List<TestHistoryVo> vos = new LinkedList<>();

        for (TestHistory po: pos) {
			TestHistoryVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TestHistoryVo genVo(TestHistory po) {
		TestHistoryVo vo = new TestHistoryVo();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}

}

