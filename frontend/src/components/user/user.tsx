import axios from "axios";
import React, { useState, useEffect } from "react";
import List from "@mui/material/List";
import TextField from "@mui/material/ListItem";

interface UserDetails {
  firstname: string;
  lastname: string;
  email: string;
  lastpaid: string;
  remaining: string;
  totalpaid: string;
  service: string;
}

const User: React.FC = () => {
  const [userDetail, setUserDetail] = useState<UserDetails>();

  const getUser = () => {
    axios
      .get("http://localhost:8080/user/ashishheda766@gmail.com")
      .then((response: any) => {
        setUserDetail(response.data);
      })
      .catch((e: Error) => {
        console.log(e);
      });
  };

  useEffect(() => {
    getUser();
  }, []);

  return (
    <List
      dense
      sx={{ width: "100%", maxWidth: 360, bgcolor: "background.paper" }}
    >
      <TextField> {userDetail?.firstname} </TextField>
      <TextField> {userDetail?.lastname} </TextField>
      <TextField> {userDetail?.email} </TextField>
      <TextField> {userDetail?.lastpaid} </TextField>
      <TextField> {userDetail?.remaining} </TextField>
      <TextField> {userDetail?.totalpaid} </TextField>
      <TextField> {userDetail?.service} </TextField>
    </List>
  );
};

export default User;
