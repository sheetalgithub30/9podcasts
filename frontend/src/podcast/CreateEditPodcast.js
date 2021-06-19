import { LeftOutlined } from "@ant-design/icons";
import { Form, Modal, Button, Input, Select, Checkbox } from "antd";
import React, { useState } from "react";

const { Option } = Select;

const CreateEditPodcast = () => {
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [title, setTitle] = useState();
  const [description, setDescription] = useState(null);
  const [domain, setDomain] = useState(null);
  const [lang, setLang] = useState("en");
  const [isExplicit, setIsExplicit] = useState(null);

  const showModal = () => {
    setIsModalVisible(true);
  };

  const handleOk = () => {
    console.log(title, description, domain, lang, isExplicit);
    // http request to server
    setIsModalVisible(false);
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };
  const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 12 },
  };
  const validateMessages = {
    required: `Label is required!`,
  };

  const onFinish = (values) => {
    console.log(values);
  };
  return (
    <>
      <Button type="primary" onClick={showModal}>
        Create a new podcast
      </Button>
      <Modal
        title="New Podcast"
        visible={isModalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
        width={650}
        height={1000}
      >
        <Form
          {...layout}
          name="podcast-form"
          onFinish={onFinish}
          validateMessages={validateMessages}
          align={LeftOutlined}
        >
          <h3>Podcast Details</h3>
          <Form.Item label="Title">
            <Input
              placeholder="A very nice podcast show title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
          </Form.Item>
          <Form.Item label="Description">
            <Input.TextArea
              placeholder="Tell us about your podcast. What makes it different?"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
          </Form.Item>
          <Form.Item label="Podcast Domain">
            <Input
              placeholder="example.com"
              value={domain}
              onChange={(e) => setDomain(e.target.value)}
            />
          </Form.Item>
          <Form.Item label="Language">
            <Input.Group compact>
              <Form.Item name={"language"} noStyle>
                <Select
                  placeholder="Language"
                  value={lang}
                  onChange={(val) => setLang(val)}
                >
                  <Option value="en">English</Option>
                  <Option value="hi">Hindi</Option>
                </Select>
              </Form.Item>
            </Input.Group>
          </Form.Item>
          <Checkbox
            value={isExplicit}
            onChange={(e) => {
              setIsExplicit(e.target.checked);
            }}
          >
            This episode includes explicit content.{" "}
          </Checkbox>
        </Form>
      </Modal>
    </>
  );
};

export default CreateEditPodcast;
