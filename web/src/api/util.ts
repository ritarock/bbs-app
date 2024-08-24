export const get = async <T>(path: string): Promise<T> => {
  return await fetch(path, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => res.json());
};

export const post = async <T>(path: string, data: T): Promise<T> => {
  return await fetch(path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  }).then((res) => res.json());
};
