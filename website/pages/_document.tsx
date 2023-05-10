import Link from "next/link";
import Image from "next/image";
import { Html, Head, Main, NextScript } from "next/document";
import { nav } from "../lib/nav";

export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <body>
        <header className="sticky top-0 z-40 w-full border-b shadow-sm flex">
          <div className="w-full flex h-12 px-auto bg-black text-white">
            <div className="container flex mx-auto">
              <SiteHeader />
            </div>
          </div>
        </header>
        <main className="flex-1 h-full bg-white text-black">
          <div className="container flex mx-auto">
            <Main />
          </div>
        </main>
        <footer>
          <SiteFooter />
        </footer>
        <NextScript />
      </body>
    </Html>
  );
}

const SiteHeader = () => {
  return (
    <>
      <>
        <div className="mr-4 flex">
          <Link href="/" className="mr-6 flex items-center space-x-2">
            <Image
              src="https://raw.githubusercontent.com/grafana/dataplane/main/docs/img/logo.svg"
              alt="dataplane"
              width={20}
              height={20}
            ></Image>
            <span className="inline-block font-bold">Grafana Data Plane</span>
          </Link>
          <nav className="flex items-center space-x-6 text-sm font-medium">
            {nav.left.map((n) => (
              <Link
                key={n.link}
                href={n.link}
                target={n.external ? "_blank" : "_self"}
              >
                {n.title}
              </Link>
            ))}
          </nav>
        </div>
        <div className="w-auto flex-1">{/* <>search goes here</> */}</div>
        <nav className="flex items-center space-x-6 text-sm font-medium">
          {nav.right.map((n) => (
            <Link
              key={n.link}
              href={n.link}
              target={n.external ? "_blank" : "_self"}
            >
              {n.title}
            </Link>
          ))}
        </nav>
      </>
    </>
  );
};

const SiteFooter = () => {
  return (
    <div className="mt-0 p-2 text-center">Copyright 2023 © Grafana Labs</div>
  );
};
