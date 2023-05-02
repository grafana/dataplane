import { tests } from "./tests";

let grafanaVersion = "unknown";

describe("panel e2e", () => {
  tests.forEach((test) => {
    const dashboardUrl = "http://localhost:3000/d/panels-e2e/panels-e2e";
    const dashboardTile = "Panels E2E";
    const tablePanelTitle = "Table Panel";
    describe(test.description, () => {
      beforeEach(() => {
        cy.visit(dashboardUrl, { qs: { "var-rawFrameUrl": test.url } });
        cy.contains(dashboardTile);
        readAndUpdateGrafanaVersion();
      });
      it("check content of the table panel", () => {
        if (test.tablePanelChecks?.length > 0) {
          checkTablePanelFullContent(
            tablePanelTitle,
            test.tablePanelChecks.join("")
          );
        }
      });
    });
  });
});

const readAndUpdateGrafanaVersion = () => {
  cy.window().then(
    (win) =>
      (grafanaVersion =
        win["grafanaBootData"]["settings"]["buildInfo"]["version"])
  );
};

const checkTablePanelFullContent = (panelTitle, textChecks: string) => {
  let grafanaMajor = +grafanaVersion.split(".")[0];
  if (grafanaMajor < 10) {
    cy.get(`section[aria-label="${panelTitle} panel"]`).contains(textChecks);
  } else {
    cy.get(
      `div[data-testid="data-testid Panel header ${panelTitle}"]`
    ).contains(textChecks);
  }
};
