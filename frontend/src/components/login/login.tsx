import {
  makeStyles,
  Container,
  Typography,
  TextField,
  Button,
} from "@material-ui/core";
import { useForm } from "react-hook-form";
import http from "../../http-common";

interface IFormInput {
  email: string;
  password: string;
}

const useStyles = makeStyles((theme) => ({
  heading: {
    textAlign: "center",
    margin: theme.spacing(1, 0, 4),
  },
  submitButton: {
    marginTop: theme.spacing(4),
  },
}));

const onSubmit = (data: IFormInput) => {
  http
    .post<IFormInput>("/login", data)
    .then((response: any) => {
      console.log(response.data);
    })
    .catch((e: Error) => {
      console.log(e);
    });
};

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<IFormInput>();

  const { heading, submitButton } = useStyles();

  return (
    <Container maxWidth="xs">
      {/* Name on the top of login form */}
      <Typography className={heading} variant="h3">
        Please Login
      </Typography>

      <form onSubmit={handleSubmit(onSubmit)} noValidate>
        <TextField
          {...register("email")}
          variant="outlined"
          margin="normal"
          label="Email"
          helperText={errors.email?.message}
          error={!!errors.email?.message}
          fullWidth
          required
        />
        <TextField
          {...register("password")}
          variant="outlined"
          margin="normal"
          label="Password"
          type="password"
          helperText={errors.email?.message}
          error={!!errors.email?.message}
          fullWidth
          required
        />
        <Button
          type="submit"
          fullWidth
          variant="contained"
          color="primary"
          className={submitButton}
        >
          Login
        </Button>
      </form>
    </Container>
  );
};

export default Login;
