import React from "react";
import { Button, Card, Input, Space } from "antd";

const Login = () => {
  return (
    <Card style={{ textAlign: "center" }}>
      <Space direction="vertical" style={{ width: "100%" }} size="middle">
        <img src="https://apnishiksha.com/logo.png" alt="logo" height="56px" />
        <Input placeholder="Email" size="large" />
        <Input placeholder="Password" size="large" />
        <Button type="primary">Login</Button>
      </Space>
    </Card>
  );
};

export default Login;
