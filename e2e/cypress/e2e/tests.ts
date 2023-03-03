export const tests: {
  description: string;
  url: string;
  tablePanelChecks?: string[];
}[] = [
  {
    description: "numeric long 0.1 two-items-by-dimension",
    url: `https://gist.githubusercontent.com/yesoreyeram/1e749c1cf927f353e7e0edd2ff047396/raw/df1088f5f93874571cf08ffc580f02becfcc61d1/numeric-sample.json`,
    tablePanelChecks: ["avgSlothCount", "city", "4", "LGA", "7.50", "MIA"],
  },
];
