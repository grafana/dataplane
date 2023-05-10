import { useRouter } from "next/router";
import { getAllPosts, getPostBySlug, markdownToHtml } from "./../../lib/api";
import markdownStyles from "./markdown-styles.module.css";

type Author = {
  name: string;
  picture: string;
};

type PostType = {
  slug: string;
  title: string;
  date: string;
  coverImage: string;
  author: Author;
  excerpt: string;
  ogImage: {
    url: string;
  };
  content: string;
};

type Props = {
  post: PostType;
  morePosts: PostType[];
  preview?: boolean;
};

export default function Page(props: Props) {
  const router = useRouter();
  return (
    <div>
      <div className="w-full mx-auto">
        <div
          className={markdownStyles["markdown"]}
          dangerouslySetInnerHTML={{ __html: props.post.content }}
        />
      </div>
    </div>
  );
}

type Params = {
  params: {
    slug: string;
  };
};

export async function getStaticProps({ params }: Params) {
  const post = getPostBySlug(params.slug, [
    "title",
    "date",
    "slug",
    "author",
    "content",
    "ogImage",
    "coverImage",
  ]);
  const content = await markdownToHtml(post.content || "");

  return {
    props: {
      post: {
        ...post,
        content,
      },
    },
  };
}

export async function getStaticPaths() {
  const posts = getAllPosts(["slug"]);

  return {
    paths: posts.map((post) => {
      return {
        params: {
          slug: post.slug,
        },
      };
    }),
    fallback: false,
  };
}
