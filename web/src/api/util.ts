const get = async <T>(path: string): Promise<T> => {
  return fetch(path).then<T>((res) => res.json());
};

const post = async <T>(path: string, data: T): Promise<T> => {
  const res = await fetch(path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  return res.json();
};

export { get, post };
