import React, { ReactNode, useEffect, useState } from "react";

interface BoxProps {
  header?: ReactNode;
  content: ReactNode;
}

const getNow = () => {
  const date = new Date();
  const yyyy = date.getFullYear();
  const mm = ("0" + (date.getMonth())).slice(-2);
  const dd = ("0" + (date.getDate())).slice(-2);
  const hour = ("0" + (date.getHours())).slice(-2);
  const minute = ("0" + (date.getMinutes())).slice(-2);

  return `${yyyy}-${mm}-${dd} ${hour}:${minute}`;
};

const Box: React.FC<BoxProps> = ({ header, content }) => {
  const [now, setNow] = useState("");

  useEffect(() => {
    setNow(getNow());
  }, []);

  return (
    <>
      <div className="box-border h-vh w-5/6 border-2 border-slate-500 mx-auto">
        <div className="flex items-center justify-between bg-blue-100">
          <p className="pl-1">
            {now}
          </p>
          <div className="text-right pr-1">{header}</div>
        </div>
        <div>
          {content}
        </div>
      </div>
    </>
  );
};

export default Box;
