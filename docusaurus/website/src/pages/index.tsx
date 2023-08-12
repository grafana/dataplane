import React from "react";
import clsx from "clsx";
import Link from "@docusaurus/Link";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Layout from "@theme/Layout";
import styles from "./index.module.css";
import DOMPurify from 'dompurify';

// function HomepageHeader() {
//   const { siteConfig } = useDocusaurusContext();
//   return (
//     <header className={clsx("hero hero--primary", styles.heroBanner)}>
//       <div className={styles.heroImageContainer}>
//         <GrafanaLogo style={{ color: "rgba(255, 255, 255, 0.2)" }} />
//       </div>
//       <div className={clsx("container", styles.heroContent)}>
//         <h1 className="hero__title">{siteConfig.title}</h1>
//         <p className="hero__subtitle">{siteConfig.tagline}</p>
//         <div className={styles.buttons}>
//           <Link
//             className="button button--primary button--lg button--outline"
//             to="/contract"
//           >
//             Get Started
//           </Link>
//         </div>
//       </div>
//     </header>
//   );
// }

// export default function Home(): JSX.Element {
//   const { siteConfig } = useDocusaurusContext();
//   return (
//     <Layout title={`${siteConfig.title}`} description={siteConfig.tagline}>
//       <HomepageHeader />
//     </Layout>
//   );
// }

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header
      className={clsx(
        "container margin-top--lg margin-bottom--md",
        styles.heroContent
      )}
    >
      <div className={clsx(styles.heroBanner)}>
        <div className={styles.heroBannerWrapper}>
          <h1
            className={clsx(
              "text--normal padding-left--sm",
              styles.heroSubtitle
            )}
          >
            {siteConfig.title}{" "}
            <span style={{ fontWeight: "normal" }}></span>
          </h1>
          <p className="margin-left--md text--bold">{siteConfig.tagline}</p>
          <Link
            className="margin-left--md button button--primary button--lg"
            to="/contract"
          >
            Get Started
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title}`}
      description="Set up a Grafana plugin environment effortlessly."
    >
      <HomepageHeader />
      <div>
        <h2>What's New</h2>
        {siteConfig.customFields.newItems.slice(0, 5).map((item) => (
          <div>
            <h3>{item.title} ({item.dateString})</h3>
            <div dangerouslySetInnerHTML={{__html: DOMPurify.sanitize(item.content, {USE_PROFILES: {html: true}})}}></div>
          </div>
        ))}
      </div>
    </Layout>
  );
}
