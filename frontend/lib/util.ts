export const toDateFormat = (t: Date) => {
  const dt = new Date(t);

  return dt.getFullYear().toString() + "-" + ("0" + dt.getMonth()).slice(-2) +
    "-" + ("0" + dt.getDate()).slice(-2) + " " +
    ("0" + dt.getHours()).slice(-2) + ":" + ("0" + dt.getMinutes()).slice(-2);
};
