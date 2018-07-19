package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.utils.BeanUtilEx;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class HistoryServiceImpl extends BaseServiceImpl implements HistoryService {

	@Override
	public List<TstHistory> listByOrg(Integer orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstHistory.class);
//
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		dc.createAlias("project", "project");
//		dc.add(Restrictions.eq("project.orgId", orgId));
//
//		dc.addOrder(Order.desc("createTime"));
//
//		Page page = findPage(dc, 0, 30);
//
//		return page.getItems();

		return null;
	}

	@Override
	public List<TstHistory> listByProject(Integer projectId, TstProject.ProjectType projectType) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstHistory.class);
//
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		if (projectType.equals(TestProject.ProjectType.project.toString())) {
//			dc.add(Restrictions.eq("projectId", projectId));
//		} else {
//			dc.createAlias("project", "project");
//			dc.add(Restrictions.eq("project.parentId", projectId));
//		}
//
//		dc.addOrder(Order.desc("createTime"));
//
//		Page page = findPage(dc, 0, 30);
//
//		return page.getItems();

		return null;
	}

    @Override
    public TstHistory getById(Integer id) {
//		TstHistory po = (TstHistory) get(TstHistory.class, id);
//		TstHistory vo = genVo(po);
//
//        return vo;

		return null;
    }
    @Override
    public TstHistory create(Integer projectId, TstUser optUser, String action,
                              String entityType, Integer entityId, String name) {
        TstHistory history = new TstHistory();

//        history.setTitle("用户" + StringUtil.highlightDict(optUser.getName())
//                + action + entityType.name + StringUtil.highlightDict(name));
//        history.setProjectId(projectId);
//        history.setEntityId(entityId);
//        history.setEntityType(entityType);
//        history.setUserId(optUser.getId());
//        saveOrUpdate(history);

        return history;
    }

	@Override
	public Map<String, List<TstHistory>> genVosByDate(List<TstHistory> historyPos) {
		Map<String, List<TstHistory>> map = new LinkedHashMap();
//		for(TstHistory his: historyPos) {
//            Date createDate = his.getCreateTime();
//            String date = DateUtils.FormatDate(createDate, "yyyy-MM-dd");
//            if (!map.containsKey(date)) {
//                map.put(date, new LinkedList());
//            }
//            map.get(date).add(genVo(his));
//		}
		return map;
	}

	@Override
	public List<TstHistory> genVos(List<TstHistory> pos) {
        List<TstHistory> vos = new LinkedList();

        for (TstHistory po: pos) {
			TstHistory vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TstHistory genVo(TstHistory po) {
		TstHistory vo = new TstHistory();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}

}

