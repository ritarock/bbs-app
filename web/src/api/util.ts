export const get = async <T>(path: string): Promise<T> => {
  const token = getToken();
  return await fetch(path, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token,
    },
  }).then((res) => res.json());
};

export const post = async <T>(path: string, data: T): Promise<T> => {
  const token = getToken();
  return await fetch(path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token,
    },
    body: JSON.stringify(data),
  }).then((res) => res.json());
};

export const postNoToken = async <T>(path: string, data: T) => {
  return await fetch(path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  }).then((res) => res.json());
};

const getToken = (): string => {
  const value = document.cookie;
  return value.split("=")[1];
};
