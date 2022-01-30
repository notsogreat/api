import axios from "axios";
import React, { useState, useEffect } from "react";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItem";

interface Users {
  firstname: string;
  lastname: string;
}

const ListUser: React.FC = () => {
  const [users, setUsers] = useState<Array<Users>>([]);

  const getUser = () => {
    axios
      .get("http://localhost:8080/users")
      .then((response: any) => {
        setUsers(response.data);
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
      {users.map((user, index) => (
        <ListItemButton alignItems="flex-start" key={index}>
          {user.firstname} {user.lastname}
        </ListItemButton>
      ))}
    </List>
  );
};

export default ListUser;
