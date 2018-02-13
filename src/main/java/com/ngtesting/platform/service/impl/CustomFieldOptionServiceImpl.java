package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestCustomFieldOption;
import com.ngtesting.platform.service.CustomFieldOptionService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.CustomFieldOptionVo;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CustomFieldOptionServiceImpl extends BaseServiceImpl implements CustomFieldOptionService {

    @Override
    public List<CustomFieldOptionVo> genVos(List<TestCustomFieldOption> pos) {
        List<CustomFieldOptionVo> vos = new LinkedList<>();

        for (TestCustomFieldOption po : pos) {
            CustomFieldOptionVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public CustomFieldOptionVo genVo(TestCustomFieldOption po) {
        CustomFieldOptionVo vo = new CustomFieldOptionVo();
        BeanUtilEx.copyProperties(vo, po);

        return vo;
    }
}
