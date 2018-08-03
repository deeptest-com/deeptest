package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CustomFieldOptionDao;
import com.ngtesting.platform.model.TstCustomFieldOption;
import com.ngtesting.platform.service.CustomFieldOptionService;
import com.ngtesting.platform.utils.BeanUtilEx;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CustomFieldOptionServiceImpl extends BaseServiceImpl implements CustomFieldOptionService {
    @Autowired
    CustomFieldOptionDao customFieldOptionDao;

    @Override
    public List<TstCustomFieldOption> listVos(Integer fieldId) {
        List ls = customFieldOptionDao.listByField(fieldId);

        List<TstCustomFieldOption> vos = genVos(ls);
        return vos;
    }

    @Override
    public TstCustomFieldOption save(TstCustomFieldOption vo) {
        if (vo.getId() == null) {
            Integer maxOrder = customFieldOptionDao.getMaxOrder(vo.getFieldId());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            customFieldOptionDao.save(vo);
        } else {
            customFieldOptionDao.update(vo);
        }

        return vo;
    }

    @Override
    public boolean delete(Integer id) {
//        getDao().delete(id);
        return true;
    }

    @Override
    public boolean changeOrderPers(Integer id, String act, Integer fieldId) {
//TstCustomFieldOption

        return true;
    }

    @Override
    public List<TstCustomFieldOption> genVos(List<TstCustomFieldOption> pos) {
//        List<TstCustomFieldOption> vos = new LinkedList<>();
//
//        for (TestCustomFieldOption po : pos) {
//            TstCustomFieldOption vo = genVo(po);
//            vos.add(vo);
//        }
//        return vos;

        return null;
    }

    @Override
    public TstCustomFieldOption genVo(TstCustomFieldOption po) {
        TstCustomFieldOption vo = new TstCustomFieldOption();
        BeanUtilEx.copyProperties(po, vo);

        return vo;
    }
}
