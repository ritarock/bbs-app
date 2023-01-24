import Link from "next/link";
import { Topic } from "../../gen/v1/ts/topic";

export const TopicList = ({ topics }: { topics: Topic[] }) => {
  return (
    <>
      <div>
        <ul>
          {topics.map((topic) => (
            <li key={topic.id}>
              <Link href={`/topics/${topic.id}`}>
                {topic.name}
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </>
  );
};
