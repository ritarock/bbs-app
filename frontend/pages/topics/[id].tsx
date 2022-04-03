import { GetServerSideProps } from "next";
import { Comment } from "../../interfaces";
export default function Topic(
  { data }: { data: { code: number; comments: Comment[] } },
) {
  return (
    <>
      <div>
        {data.comments.map((e) => <div key={e.id}>{e.body}</div>)}
      </div>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ params }) => {
  const url = `http://localhost:8080/backend/api/topics/${params.id}/comments`;
  const res = await fetch(url);
  const data = await res.json();
  console.log(data);

  return {
    props: { data },
  };
};
