package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstCase;
import org.apache.poi.ss.usermodel.Sheet;
import org.apache.poi.xssf.usermodel.XSSFCellStyle;

import java.util.List;

public interface CaseExportService extends BaseService {

    String export(Integer projectId);

	Integer writeHeader(Sheet sheet, Integer rowCount, XSSFCellStyle cellStyle);

	Integer writeTestCase(TstCase testCase, Sheet sheet,  Integer rowCount, XSSFCellStyle cellStyle, XSSFCellStyle optStyle);

    List<TstCase> sortParentAndChild(List<TstCase> entities);

    void setLevel(List<TstCase> entities, Integer level);
}
