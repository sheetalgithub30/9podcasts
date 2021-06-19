import { Row, Col } from "antd";
import React from "react";
import CreateEditPodcast from "../podcast/CreateEditPodcast";

const NoPodcast = () => {
  return (
    <Row justify="center" align="middle" style={{ height: "500px" }}>
      <Col align="center">
        <p>You do not have any podcasts.</p>
        <CreateEditPodcast />
      </Col>
    </Row>
  );
};

export default NoPodcast;
