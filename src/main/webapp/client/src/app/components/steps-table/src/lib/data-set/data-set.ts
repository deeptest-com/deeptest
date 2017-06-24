import { Row } from './row';
import { Column } from './column';

export class DataSet {

  newRow: Row;

  protected data: Array<any> = [];
  protected columns: Array<Column> = [];
  protected rows: Array<Row> = [];

  constructor(data: Array<any> = [], protected columnSettings: Object) {
    this.createColumns(columnSettings);
    this.setData(data);

    this.createNewRow();
  }

  setData(data: Array<any>) {
    this.data = data;
    this.createRows();
  }

  getColumns(): Array<Column> {
    return this.columns;
  }

  getRows(): Array<Row> {
    return this.rows;
  }

  findRowByData(data: any): Row {
    return this.rows.find((row: Row) => row.getData() === data);
  }

  createNewRow() {
    this.newRow = new Row(-1, {}, this);
    this.newRow.isInEditing = true;
  }

  /**
   * Create columns by mapping from the settings
   * @param settings
   * @private
   */
  createColumns(settings: any) {
    for (const id in settings) {
      if (settings.hasOwnProperty(id)) {
        this.columns.push(new Column(id, settings[id], this));
      }
    }
  }

  /**
   * Create rows based on current data prepared in data source
   * @private
   */
  createRows() {
    this.rows = [];
    this.data.forEach((el, index) => {
      this.rows.push(new Row(index, el, this));
    });
  }
}
