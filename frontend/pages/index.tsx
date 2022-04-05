import { GetServerSideProps } from "next";
import Link from "next/link";
import { useRouter } from "next/router";
import Form from "../components/form";
import { Topic } from "../interfaces";
import { toDateFormat } from "../lib/util";

export default function Home(
  { data }: { data: { code: number; topics: Topic[] } },
) {
  const router = useRouter();
  return (
    <>
      <div>
        <ul>
          {data.topics.map((e) => (
            <li key={e.id}>
              <Link href={`/topics/${e.id}`}>
                {e.name}
              </Link>
              : {toDateFormat(e.created_at)}
            </li>
          ))}
        </ul>
      </div>
      <div>
        <Form
          postUrl={"http://localhost:8080/backend/api/topics"}
        />
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
