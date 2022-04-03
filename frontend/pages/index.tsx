import { GetServerSideProps } from "next";
import Link from "next/link";
import { Topic } from "../interfaces";
import { toTitleFormat } from "../lib/util";

export default function Home(
  { data }: { data: { code: number; topics: Topic[] } },
) {
  return (
    <>
      <div>
        <ul>
          {data.topics.map((e) => (
            <li key={e.id}>
              <Link href={`/topics/${e.id}`}>
                {e.name}
              </Link>
              : {toTitleFormat(e.created_at)}
            </li>
          ))}
        </ul>
      </div>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async () => {
  const url = "http://localhost:8080/backend/api/topics";
  const res = await fetch(url);
  const data = await res.json();

  return {
    props: { data },
  };
};
