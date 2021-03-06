var React = require('react');
var Page = require('../components/Page');
var PageHeader = require('../components/PageHeader');
var EntityBreadcrumb = require('../components/EntityBreadcrumb');
var EntityButtonGroup = require('../components/EntityButtonGroup');
var Bootstrap = require('react-bootstrap');
var Row = Bootstrap.Row;
var Col = Bootstrap.Col;
var Breadcrumb = Bootstrap.Breadcrumb;
var ButtonGroup = Bootstrap.ButtonGroup;
var Button = Bootstrap.Button;
var Table = Bootstrap.Table;
var Navbar = Bootstrap.Navbar;
var FormGroup = Bootstrap.FormGroup;
var FormControl = Bootstrap.FormControl;

var layer0 = require('../utils/layer0');

var TaskDashboardContainer = React.createClass({
  getInitialState: function() {
    return {
      isLoading: true,
      taskRows: []
    }
  },
  render: function () {
    environmentID = this.props.routeParams.environmentID

    return (
      <Page>

        <Row>
          <Col lg={12}>
            <PageHeader text='Dashboard' subtext={environmentID} />
            <EntityBreadcrumb environmentID={environmentID} tasks />
          </Col>
        </Row>

        <Row>
          <Col lg={12}>
            <EntityButtonGroup environmentID={environmentID} tasks/>
          </Col>
        </Row>

        <Row>
          <Col lg={12}>
            <Table responsive striped condensed hover>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Name</th>
                  <th>Deploy</th>
                  <th>Copies</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>some id</td>
                  <td>some name</td>
                  <td>some deploy</td>
                  <td>1</td>
                </tr>
                <tr>
                  <td>some id</td>
                  <td>some name</td>
                  <td>some deploy</td>
                  <td>2</td>
                </tr>
              </tbody>
            </Table>
          </Col>
        </Row>    

      </Page>
    )
  }
});

module.exports = TaskDashboardContainer;
